package idle

import (
	"sync"
)

// AlibabaIdleTenderAftersaleOrderGetModule 结构体
type AlibabaIdleTenderAftersaleOrderGetModule struct {
	// 申请服务类型
	ApplyType string `json:"apply_type,omitempty" xml:"apply_type,omitempty"`
	// 理赔方案类型
	ServicePlanType string `json:"service_plan_type,omitempty" xml:"service_plan_type,omitempty"`
	// 理赔方案信息
	ServicePlanInfo string `json:"service_plan_info,omitempty" xml:"service_plan_info,omitempty"`
	// 正向订单号
	MainOrderId string `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 申请信息
	ApplyInfo string `json:"apply_info,omitempty" xml:"apply_info,omitempty"`
	// 售后服务CODE
	ServiceBizCode string `json:"service_biz_code,omitempty" xml:"service_biz_code,omitempty"`
	// 售后服务场景类型
	ServiceSceneType string `json:"service_scene_type,omitempty" xml:"service_scene_type,omitempty"`
	// 帮卖完结方式
	MainOrderSettleType string `json:"main_order_settle_type,omitempty" xml:"main_order_settle_type,omitempty"`
	// 售后申请单主状态
	MainStatus int64 `json:"main_status,omitempty" xml:"main_status,omitempty"`
	// 扩展信息
	Attributes *Attributes `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 售后申请单子状态
	SubStatus int64 `json:"sub_status,omitempty" xml:"sub_status,omitempty"`
}

var poolAlibabaIdleTenderAftersaleOrderGetModule = sync.Pool{
	New: func() any {
		return new(AlibabaIdleTenderAftersaleOrderGetModule)
	},
}

// GetAlibabaIdleTenderAftersaleOrderGetModule() 从对象池中获取AlibabaIdleTenderAftersaleOrderGetModule
func GetAlibabaIdleTenderAftersaleOrderGetModule() *AlibabaIdleTenderAftersaleOrderGetModule {
	return poolAlibabaIdleTenderAftersaleOrderGetModule.Get().(*AlibabaIdleTenderAftersaleOrderGetModule)
}

// ReleaseAlibabaIdleTenderAftersaleOrderGetModule 释放AlibabaIdleTenderAftersaleOrderGetModule
func ReleaseAlibabaIdleTenderAftersaleOrderGetModule(v *AlibabaIdleTenderAftersaleOrderGetModule) {
	v.ApplyType = ""
	v.ServicePlanType = ""
	v.ServicePlanInfo = ""
	v.MainOrderId = ""
	v.ApplyInfo = ""
	v.ServiceBizCode = ""
	v.ServiceSceneType = ""
	v.MainOrderSettleType = ""
	v.MainStatus = 0
	v.Attributes = nil
	v.SubStatus = 0
	poolAlibabaIdleTenderAftersaleOrderGetModule.Put(v)
}
