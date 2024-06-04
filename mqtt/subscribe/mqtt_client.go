package subscribe

import (
	"os"
	"time"

	"project/initialize"
	config "project/mqtt"

	"project/mqtt/publish"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-basic/uuid"
	"github.com/panjf2000/ants"
	"github.com/sirupsen/logrus"
)

var SubscribeMqttClient mqtt.Client
var TelemetryMessagesChan chan map[string]interface{}

func SubscribeInit() {

	//实例限流客户端
	initialize.NewAutomateLimiter()
	// 创建mqtt客户端
	subscribeMqttClient()
	// 创建消息队列
	telemetryMessagesChan()
	// 订阅attribute消息
	SubscribeAttribute()
	// 订阅event消息
	SubscribeEvent()
	//订阅telemetry消息
	err := SubscribeTelemetry()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	// 订阅在线离线消息
	SubscribeDeviceStatus()

	//网关订阅主题
	GatewaySubscribeTopic()

	// 订阅设备命令消息
	SubscribeCommand()
}

func subscribeMqttClient() {
	// 初始化配置
	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.MqttConfig.Broker)
	opts.SetUsername(config.MqttConfig.User)
	opts.SetPassword(config.MqttConfig.Pass)
	opts.SetClientID("thingspanel-go-sub-" + uuid.New()[0:8])
	// 干净会话
	opts.SetCleanSession(true)
	// 恢复客户端订阅，需要broker支持
	opts.SetResumeSubs(true)
	// 自动重连
	opts.SetAutoReconnect(true)
	opts.SetConnectRetryInterval(5 * time.Second)
	opts.SetMaxReconnectInterval(200 * time.Second)
	// 消息顺序
	opts.SetOrderMatters(false)
	opts.SetOnConnectHandler(func(c mqtt.Client) {
		logrus.Println("mqtt connect success")
	})

	SubscribeMqttClient = mqtt.NewClient(opts)
	// 等待连接成功，失败重新连接
	for {
		if token := SubscribeMqttClient.Connect(); token.Wait() && token.Error() != nil {
			logrus.Error("MQTT Broker 1 连接失败:", token.Error())
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

}

// 创建消息队列
func telemetryMessagesChan() {
	TelemetryMessagesChan = make(chan map[string]interface{}, config.MqttConfig.ChannelBufferSize)
	writeWorkers := config.MqttConfig.WriteWorkers
	for i := 0; i < writeWorkers; i++ {
		go MessagesChanHandler(TelemetryMessagesChan)
	}
}

// 订阅telemetry消息
func SubscribeTelemetry() error {

	p, err := ants.NewPool(config.MqttConfig.Telemetry.PoolSize)
	if err != nil {
		return err
	}
	deviceTelemetryMessageHandler := func(c mqtt.Client, d mqtt.Message) {
		err = p.Submit(func() {
			// 处理消息
			TelemetryMessages(d.Payload(), d.Topic())
		})
		if err != nil {
			logrus.Error(err)
		}
	}

	topic := config.MqttConfig.Telemetry.SubscribeTopic
	qos := byte(config.MqttConfig.Telemetry.QoS)

	if token := SubscribeMqttClient.Subscribe(topic, qos, deviceTelemetryMessageHandler); token.Wait() && token.Error() != nil {
		logrus.Error(token.Error())
		os.Exit(1)
	}
	return nil
}

// 订阅attribute消息，暂不需要线程池，不需要消息队列
func SubscribeAttribute() {
	// 订阅attribute消息
	deviceAttributeHandler := func(c mqtt.Client, d mqtt.Message) {
		// 处理消息
		logrus.Debug("attribute message:", string(d.Payload()))
		deviceNumber, messageId, err := DeviceAttributeReport(d.Payload(), d.Topic())
		logrus.Debug("响应设备属性上报", deviceNumber, err)
		if err != nil {
			logrus.Error(err)
		}
		if deviceNumber != "" && messageId != "" {
			// 响应设备属性上报
			publish.PublishAttributeResponseMessage(deviceNumber, messageId, err)
		}
	}
	topic := config.MqttConfig.Attributes.SubscribeTopic
	logrus.Debug("subscribe topic:", topic)
	qos := byte(config.MqttConfig.Attributes.QoS)
	if token := SubscribeMqttClient.Subscribe(topic, qos, deviceAttributeHandler); token.Wait() && token.Error() != nil {
		logrus.Error(token.Error())
		os.Exit(1)
	}
}

// 订阅command消息，暂不需要线程池，不需要消息队列
func SubscribeCommand() {
	// 订阅command消息
	deviceCommandHandler := func(c mqtt.Client, d mqtt.Message) {
		// 处理消息
		messageID, err := DeviceCommand(d.Payload(), d.Topic())
		logrus.Debug("设备命令响应上报", messageID, err)
		if err != nil || messageID == "" {
			logrus.Debug("设备命令响应上报失败", messageID, err)
			logrus.Error(err)
		}
	}
	topic := config.MqttConfig.Commands.SubscribeTopic
	qos := byte(config.MqttConfig.Commands.QoS)
	if token := SubscribeMqttClient.Subscribe(topic, qos, deviceCommandHandler); token.Wait() && token.Error() != nil {
		logrus.Error(token.Error())
		os.Exit(1)
	}
}

// 订阅event消息，暂不需要线程池，不需要消息队列
func SubscribeEvent() {
	// 订阅event消息
	deviceEventHandler := func(c mqtt.Client, d mqtt.Message) {
		// 处理消息
		logrus.Debug("event message:", string(d.Payload()))
		deviceNumber, messageId, method, err := DeviceEvent(d.Payload(), d.Topic())
		logrus.Debug("响应设备属性上报", deviceNumber, err)
		if err != nil {
			logrus.Error(err)
		}
		if deviceNumber != "" && messageId != "" {
			// 响应设备属性上报
			publish.PublishEventResponseMessage(deviceNumber, messageId, method, err)
		}
	}
	topic := config.MqttConfig.Events.SubscribeTopic
	qos := byte(config.MqttConfig.Events.QoS)
	if token := SubscribeMqttClient.Subscribe(topic, qos, deviceEventHandler); token.Wait() && token.Error() != nil {
		logrus.Error(token.Error())
		os.Exit(1)
	}
}

// 订阅设备上线离线消息
func SubscribeDeviceStatus() {
	// 订阅设备上线离线消息
	deviceOnlineHandler := func(c mqtt.Client, d mqtt.Message) {
		logrus.Debug("设备已上线")
		// 处理消息
		DeviceOnline(d.Payload(), d.Topic())

	}
	onlineTopic := "devices/status/+"
	onlineQos := byte(1)
	if token := SubscribeMqttClient.Subscribe(onlineTopic, onlineQos, deviceOnlineHandler); token.Wait() && token.Error() != nil {
		logrus.Error(token.Error())
		os.Exit(1)
	}
}

func SubscribeOtaUpprogress() {
	// 订阅ota升级消息
	otaUpgradeHandler := func(c mqtt.Client, d mqtt.Message) {
		// 处理消息
		logrus.Debug("ota upgrade message:", string(d.Payload()))
		OtaUpgrade(d.Payload(), d.Topic())
	}
	topic := config.MqttConfig.OTA.SubscribeTopic
	qos := byte(config.MqttConfig.OTA.QoS)
	if token := SubscribeMqttClient.Subscribe(topic, qos, otaUpgradeHandler); token.Wait() && token.Error() != nil {
		logrus.Error(token.Error())
		os.Exit(1)
	}
}
