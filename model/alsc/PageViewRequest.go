package alsc

import (
	"sync"
)

// PageViewRequest 结构体
type PageViewRequest struct {
	// 账号转换方案 当不传的时候，不做账号转换
	AccountPlan string `json:"account_plan,omitempty" xml:"account_plan,omitempty"`
	// 行为编码
	ActionCode string `json:"action_code,omitempty" xml:"action_code,omitempty"`
	// 业务场景
	BizScene string `json:"biz_scene,omitempty" xml:"biz_scene,omitempty"`
	// 渠道
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 城市id
	CityId string `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 端来源
	Client string `json:"client,omitempty" xml:"client,omitempty"`
	// 端用户类型
	ClientUserType string `json:"client_user_type,omitempty" xml:"client_user_type,omitempty"`
	// 任务集id
	CollectionId string `json:"collection_id,omitempty" xml:"collection_id,omitempty"`
	// 扩展参数
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	//  纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 风控token
	MeetResultToken string `json:"meet_result_token,omitempty" xml:"meet_result_token,omitempty"`
	// 任务id
	MissionId string `json:"mission_id,omitempty" xml:"mission_id,omitempty"`
	// 任务xid
	MissionXId string `json:"mission_x_id,omitempty" xml:"mission_x_id,omitempty"`
	// 页面来源 如果任务后台配置了，则需要填，不然任务推进不了
	PageFrom string `json:"page_from,omitempty" xml:"page_from,omitempty"`
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// shareId
	ShareId string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// 微信）风控参数
	Ua string `json:"ua,omitempty" xml:"ua,omitempty"`
	// 微信）风控参数
	UmiToken string `json:"umi_token,omitempty" xml:"umi_token,omitempty"`
	// 微信）风控参数
	UnionId string `json:"union_id,omitempty" xml:"union_id,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 浏览时间 如果任务后台配置了，则需要填，不然任务推进不了。
	ViewTime string `json:"view_time,omitempty" xml:"view_time,omitempty"`
	// 用户id类型 账号类型，0表示淘宝账号，25表示饿了么账号
	UserIdType string `json:"user_id_type,omitempty" xml:"user_id_type,omitempty"`
	// 是否是支付宝三方小程序
	AlipayThirdMiniProgram bool `json:"alipay_third_mini_program,omitempty" xml:"alipay_third_mini_program,omitempty"`
	// 是否是小程序
	MiniApp bool `json:"mini_app,omitempty" xml:"mini_app,omitempty"`
	// 是否同步
	Sync bool `json:"sync,omitempty" xml:"sync,omitempty"`
}

var poolPageViewRequest = sync.Pool{
	New: func() any {
		return new(PageViewRequest)
	},
}

// GetPageViewRequest() 从对象池中获取PageViewRequest
func GetPageViewRequest() *PageViewRequest {
	return poolPageViewRequest.Get().(*PageViewRequest)
}

// ReleasePageViewRequest 释放PageViewRequest
func ReleasePageViewRequest(v *PageViewRequest) {
	v.AccountPlan = ""
	v.ActionCode = ""
	v.BizScene = ""
	v.Channel = ""
	v.CityId = ""
	v.Client = ""
	v.ClientUserType = ""
	v.CollectionId = ""
	v.ExtInfo = ""
	v.Latitude = ""
	v.Longitude = ""
	v.MeetResultToken = ""
	v.MissionId = ""
	v.MissionXId = ""
	v.PageFrom = ""
	v.RequestId = ""
	v.ShareId = ""
	v.Ua = ""
	v.UmiToken = ""
	v.UnionId = ""
	v.UserId = ""
	v.ViewTime = ""
	v.UserIdType = ""
	v.AlipayThirdMiniProgram = false
	v.MiniApp = false
	v.Sync = false
	poolPageViewRequest.Put(v)
}
