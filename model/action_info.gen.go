// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameActionInfo = "action_info"

// ActionInfo mapped from table <action_info>
type ActionInfo struct {
	ID                string  `gorm:"column:id;primaryKey" json:"id"`
	SceneAutomationID string  `gorm:"column:scene_automation_id;not null;comment:场景联动ID（外键-关联删除）" json:"scene_automation_id"`                 // 场景联动ID（外键-关联删除）
	ActionTarget      *string `gorm:"column:action_target;comment:动作目标id设备id、场景id、告警id；如果条件是单类设备，这里为空" json:"action_target"`                  // 动作目标id设备id、场景id、告警id；如果条件是单类设备，这里为空
	ActionType        string  `gorm:"column:action_type;not null;comment:动作类型10: 单个设备11: 单类设备20: 激活场景30: 触发告警40: 服务" json:"action_type"` // 动作类型10: 单个设备11: 单类设备20: 激活场景30: 触发告警40: 服务
	ActionParamType   *string `gorm:"column:action_param_type;comment:遥测TEL属性ATTR命令CMD" json:"action_param_type"`                             // 遥测TEL属性ATTR命令CMD
	ActionParam       *string `gorm:"column:action_param;comment:动作参数动作类型为10,11是有效 标识符" json:"action_param"`                                 // 动作参数动作类型为10,11是有效 标识符
	ActionValue       *string `gorm:"column:action_value;comment:目标值" json:"action_value"`                                                    // 目标值
	Remark            *string `gorm:"column:remark" json:"remark"`
}

// TableName ActionInfo's table name
func (*ActionInfo) TableName() string {
	return TableNameActionInfo
}
