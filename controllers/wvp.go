package controllers

import (
	"ThingsPanel-Go/services"
	response "ThingsPanel-Go/utils"
	"encoding/json"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
	context2 "github.com/beego/beego/v2/server/web/context"
)

type WvpController struct {
	beego.Controller
}

// DeleteWidget 校验
type PtzControlValid struct {
	ParentId      string `json:"parent_id" alias:"ID" valid:"Required; MaxSize(36)"`
	SubDeviceAddr string `json:"sub_device_addr" alias:"ID" valid:"Required;MaxSize(36)"`
}

// 修改扩展功能
func (widgetController *WvpController) PtzControl() {
	var ptzControlValid = make(map[string]string)
	err := json.Unmarshal(widgetController.Ctx.Input.RequestBody, &ptzControlValid)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	var parentId, subDeviceAddr string
	if ptzControlValid["parent_id"] == "" {
		response.SuccessWithMessage(400, "parent_id不能为空", (*context2.Context)(widgetController.Ctx))
		return
	} else {
		parentId = ptzControlValid["parent_id"]
		delete(ptzControlValid, "parent_id")
	}
	if ptzControlValid["sub_device_addr"] == "" {
		response.SuccessWithMessage(400, "sub_device_addr不能为空", (*context2.Context)(widgetController.Ctx))
		return
	} else {
		subDeviceAddr = ptzControlValid["sub_device_addr"]
		delete(ptzControlValid, "sub_device_addr")
	}
	var wvpDeviceService services.WvpDeviceService
	err = wvpDeviceService.PtzControl(parentId, subDeviceAddr, ptzControlValid)
	if err == nil {
		// 修改成功
		response.SuccessWithMessage(200, "success", (*context2.Context)(widgetController.Ctx))
		return
	}
	// 修改失败
	response.SuccessWithMessage(400, err.Error(), (*context2.Context)(widgetController.Ctx))
}
