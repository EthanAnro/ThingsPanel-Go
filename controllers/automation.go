package controllers

import (
	gvalid "ThingsPanel-Go/initialize/validate"
	"ThingsPanel-Go/models"
	"ThingsPanel-Go/services"
	response "ThingsPanel-Go/utils"
	valid "ThingsPanel-Go/validate"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	context2 "github.com/beego/beego/v2/server/web/context"
	"github.com/bitly/go-simplejson"
)

type AutomationController struct {
	beego.Controller
}

type SymbolData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Configrules struct {
	Device       []models.Asset   `json:"device"`
	AssemblyArr  []services.Field `json:"assemblyArr"`
	ConditionArr []ConditionArr   `json:"conditionArr"`
}

type ConfigApply struct {
	Device       []models.Asset   `json:"device"`
	AssemblyArr  []services.Field `json:"assemblyArr"`
	ConditionArr []ConditionArr   `json:"conditionArr"`
}

type ConfigAll struct {
	Rules []Configrules `json:"rules"`
	Apply []ConfigApply `json:"apply"`
}

type widgetList struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type ConditionArr struct {
	ID      string       `json:"id"`
	Name    string       `json:"name"`
	Widgets []widgetList `json:"widgets"`
}

type AutoUpdate struct {
	ID         string    `json:"id"`
	BusinessID string    `json:"business_id"`
	Name       string    `json:"name"`
	Describe   string    `json:"describe"`
	Status     string    `json:"status"`
	Config     ConfigAll `json:"config"`
	Sort       int64     `json:"sort"`
	Type       int64     `json:"type"`
	Issued     string    `json:"issued"`
	CustomerID string    `json:"customer_id"`
}

// 策略列表
func (this *AutomationController) Index() {
	automationIndexValidate := valid.AutomationIndex{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &automationIndexValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(automationIndexValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(automationIndexValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var AutomationService services.AutomationService
	u, _ := AutomationService.Paginate(automationIndexValidate.BusinessId, automationIndexValidate.Limit, automationIndexValidate.Page-1)
	response.SuccessWithDetailed(200, "success", u, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

// 添加策略
func (this *AutomationController) Add() {
	automationAddValidate := valid.AutomationAdd{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &automationAddValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(automationAddValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(automationAddValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var AutomationService services.AutomationService
	f, _ := AutomationService.Add(
		automationAddValidate.BusinessID,
		automationAddValidate.Name,
		automationAddValidate.Describe,
		automationAddValidate.Status,
		automationAddValidate.Config,
		automationAddValidate.Sort,
		automationAddValidate.Type,
		automationAddValidate.Issued,
		automationAddValidate.CustomerID,
	)
	if f {
		response.SuccessWithMessage(200, "success", (*context2.Context)(this.Ctx))
		return
	}
	// 新增失败
	response.SuccessWithMessage(400, "error", (*context2.Context)(this.Ctx))
	return
}

// 编辑策略
func (this *AutomationController) Edit() {
	automationEditValidate := valid.AutomationEdit{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &automationEditValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(automationEditValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(automationEditValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var AutomationService services.AutomationService
	f := AutomationService.Edit(
		automationEditValidate.ID,
		automationEditValidate.BusinessID,
		automationEditValidate.Name,
		automationEditValidate.Describe,
		automationEditValidate.Status,
		automationEditValidate.Config,
		automationEditValidate.Sort,
		automationEditValidate.Type,
		automationEditValidate.Issued,
		automationEditValidate.CustomerID,
	)
	if f {
		// 编辑成功
		response.SuccessWithMessage(200, "success", (*context2.Context)(this.Ctx))
		return
	}
	// 编辑失败
	response.SuccessWithMessage(400, "error", (*context2.Context)(this.Ctx))
	return
}

// 删除策略
func (this *AutomationController) Delete() {
	automationDeleteValidate := valid.AutomationDelete{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &automationDeleteValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(automationDeleteValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(automationDeleteValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var AutomationService services.AutomationService
	f := AutomationService.Delete(automationDeleteValidate.ID)
	if f {
		response.SuccessWithMessage(200, "删除成功", (*context2.Context)(this.Ctx))
		return
	}
	response.SuccessWithMessage(400, "删除失败", (*context2.Context)(this.Ctx))
	return
}

// 获取获取具体某一条策略
func (this *AutomationController) GetOne() {
	automationGetValidate := valid.AutomationGet{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &automationGetValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(automationGetValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(automationGetValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(200, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var AutomationService services.AutomationService
	u, i := AutomationService.GetAutomationById(automationGetValidate.ID)
	if i != 0 {
		// 获取成功
		response.SuccessWithDetailed(200, "success", u, map[string]string{}, (*context2.Context)(this.Ctx))
		return
	}
	// 获取失败
	response.SuccessWithMessage(500, "error", (*context2.Context)(this.Ctx))
	return
}

// 状态码
func (this *AutomationController) Status() {
	d := [...]string{"每天执行", "每一分钟执行一次", "每五分钟执行一次", "每十分钟执行一次", "每一小时执行一次", "每三小时执行一次", "每六小时执行一次", "每十二小时执行一次"}
	response.SuccessWithDetailed(200, "success", d, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

// 逻辑符
func (this *AutomationController) Symbol() {
	var sd []SymbolData
	d1 := SymbolData{
		ID:   "<",
		Name: "小于",
	}
	sd = append(sd, d1)
	d2 := SymbolData{
		ID:   ">",
		Name: "大于",
	}
	sd = append(sd, d2)
	d3 := SymbolData{
		ID:   "==",
		Name: "等于",
	}
	sd = append(sd, d3)
	d4 := SymbolData{
		ID:   "<=",
		Name: "小于等于",
	}
	sd = append(sd, d4)
	d5 := SymbolData{
		ID:   ">=",
		Name: "大于等于",
	}
	sd = append(sd, d5)
	response.SuccessWithDetailed(200, "success", sd, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

//
func (this *AutomationController) Property() {
	automationPropertyValidate := valid.AutomationProperty{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &automationPropertyValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(automationPropertyValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(automationPropertyValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(200, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var DeviceService services.DeviceService
	d, _ := DeviceService.GetDeviceByID(automationPropertyValidate.BusinessID)
	response.SuccessWithDetailed(200, "success", d, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

// Show
func (this *AutomationController) Show() {
	automationShowValidate := valid.AutomationShow{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &automationShowValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(automationShowValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(automationShowValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var DeviceService services.DeviceService
	var AssetService services.AssetService
	var fd []services.Field
	d, _ := DeviceService.GetDeviceByID(automationShowValidate.Bid)
	wl := AssetService.Widget(d.Type)
	if len(wl) > 0 {
		for _, wv := range wl {
			fl := AssetService.Field(d.Type, wv.Key)
			if len(fl) > 0 {
				for _, fv := range fl {
					if fv.Type != 0 && fv.Type != 4 && fv.Type != 5 {
						fd = append(fd, fv)
					}
				}
			}
		}
	}
	response.SuccessWithDetailed(200, "success", fd, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

// Update
func (this *AutomationController) Update() {
	automationUpdateValidate := valid.AutomationUpdate{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &automationUpdateValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(automationUpdateValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(automationUpdateValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var ConditionsService services.ConditionsService
	var AssetService services.AssetService
	var DeviceService services.DeviceService
	co, _ := ConditionsService.GetConditionByID(automationUpdateValidate.ID)
	de, _ := AssetService.GetAssetData(co.BusinessID)
	res, e := simplejson.NewJson([]byte(co.Config))
	var ca []Configrules
	var ca2 []ConfigApply
	if e != nil {
		fmt.Println("解析出错", e)
	}
	if co.Type == 1 {
		rows, _ := res.Get("rules").Array()
		for _, row := range rows {
			var fd []services.Field
			var widgetValue []widgetList
			var component []ConditionArr
			ri, _ := row.(map[string]interface{})
			d, _ := DeviceService.GetDeviceByID(ri["device_id"].(string))
			wl := AssetService.Widget(d.Type)
			if len(wl) > 0 {
				for _, wv := range wl {
					fl := AssetService.Field(d.Type, wv.Key)
					if len(fl) > 0 {
						for _, fv := range fl {
							if fv.Type != 0 && fv.Type != 4 && fv.Type != 5 {
								fd = append(fd, fv)
							}
						}
					}
				}
			}
			dl, dc := DeviceService.GetDevicesByAssetID(ri["asset_id"].(string))
			if dc > 0 {
				for _, dv := range dl {
					el := AssetService.Extension()
					var n string
					if len(el) > 0 {
						for _, ev := range el {
							if ev.Key == dv.Type {
								n = ev.Name
							}
						}
					}
					wl2 := AssetService.Widget(dv.Type)
					if len(wl2) > 0 {
						for _, wv2 := range wl2 {
							i2 := widgetList{
								Name: wv2.Name,
								Key:  dv.Type + ":" + wv2.Key,
							}
							widgetValue = append(widgetValue, i2)
						}
					}
					c := ConditionArr{
						ID:      dv.ID,
						Name:    n,
						Widgets: widgetValue,
					}
					component = append(component, c)
				}
			}
			cai := Configrules{
				Device:       de,
				AssemblyArr:  fd,
				ConditionArr: component,
			}
			ca = append(ca, cai)
		}
	}

	rows2, _ := res.Get("apply").Array()
	for _, row2 := range rows2 {
		var fd2 []services.Field
		var widgetValue2 []widgetList
		var component2 []ConditionArr
		ri2, _ := row2.(map[string]interface{})
		d2, _ := DeviceService.GetDeviceByID(ri2["device_id"].(string))
		wl2 := AssetService.Widget(d2.Type)
		if len(wl2) > 0 {
			for _, wv2 := range wl2 {
				fl2 := AssetService.Field(d2.Type, wv2.Key)
				if len(fl2) > 0 {
					for _, fv2 := range fl2 {
						if fv2.Type != 0 && fv2.Type != 1 && fv2.Type != 4 && fv2.Type != 5 {
							fd2 = append(fd2, fv2)
						}
					}
				}
			}
		}
		dl2, dc2 := DeviceService.GetDevicesByAssetID(ri2["asset_id"].(string))
		if dc2 > 0 {
			for _, dv2 := range dl2 {
				el2 := AssetService.Extension()
				var n2 string
				if len(el2) > 0 {
					for _, ev2 := range el2 {
						if ev2.Key == dv2.Type {
							n2 = ev2.Name
						}
					}
				}
				wl22 := AssetService.Widget(dv2.Type)
				if len(wl22) > 0 {
					for _, wv22 := range wl22 {
						i22 := widgetList{
							Name: wv22.Name,
							Key:  dv2.Type + ":" + wv22.Key,
						}
						widgetValue2 = append(widgetValue2, i22)
					}
				}
				c2 := ConditionArr{
					ID:      dv2.ID,
					Name:    n2,
					Widgets: widgetValue2,
				}
				component2 = append(component2, c2)
			}
		}
		cai2 := ConfigApply{
			Device:       de,
			AssemblyArr:  fd2,
			ConditionArr: component2,
		}
		ca2 = append(ca2, cai2)
	}
	config := ConfigAll{
		Rules: ca,
		Apply: ca2,
	}
	resp := AutoUpdate{
		ID:         co.ID,
		BusinessID: co.BusinessID,
		Name:       co.Name,
		Describe:   co.Describe,
		Status:     co.Status,
		Config:     config,
		Sort:       co.Sort,
		Type:       co.Type,
		Issued:     co.Issued,
		CustomerID: co.CustomerID,
	}
	response.SuccessWithDetailed(200, "success", resp, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

// Instruct
func (this *AutomationController) Instruct() {
	automationInstructValidate := valid.AutomationInstruct{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &automationInstructValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(automationInstructValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(automationInstructValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var DeviceService services.DeviceService
	var AssetService services.AssetService
	var fd []services.Field
	d, _ := DeviceService.GetDeviceByID(automationInstructValidate.Bid)
	wl := AssetService.Widget(d.Type)
	if len(wl) > 0 {
		for _, wv := range wl {
			fl := AssetService.Field(d.Type, wv.Key)
			if len(fl) > 0 {
				for _, fv := range fl {
					if fv.Type != 0 && fv.Type != 1 && fv.Type != 4 && fv.Type != 5 {
						fd = append(fd, fv)
					}
				}
			}
		}
	}
	response.SuccessWithDetailed(200, "success", fd, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}