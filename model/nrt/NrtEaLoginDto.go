package nrt

import (
	"sync"
)

// NrtEaLoginDto 结构体
type NrtEaLoginDto struct {
	// 绑定类型
	ActionType string `json:"action_type,omitempty" xml:"action_type,omitempty"`
	// 企业名称
	EntName string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
	// 业务身份
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 绑定的用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 场景
	Scene string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 登录地址
	RedirectUrl string `json:"redirect_url,omitempty" xml:"redirect_url,omitempty"`
	// 被绑定用户Id
	BindSiteUserId int64 `json:"bind_site_user_id,omitempty" xml:"bind_site_user_id,omitempty"`
	// 企业id
	EntId int64 `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
	// 操作时间戳
	OperateTime int64 `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
}

var poolNrtEaLoginDto = sync.Pool{
	New: func() any {
		return new(NrtEaLoginDto)
	},
}

// GetNrtEaLoginDto() 从对象池中获取NrtEaLoginDto
func GetNrtEaLoginDto() *NrtEaLoginDto {
	return poolNrtEaLoginDto.Get().(*NrtEaLoginDto)
}

// ReleaseNrtEaLoginDto 释放NrtEaLoginDto
func ReleaseNrtEaLoginDto(v *NrtEaLoginDto) {
	v.ActionType = ""
	v.EntName = ""
	v.BizCode = ""
	v.Mobile = ""
	v.UserId = ""
	v.Scene = ""
	v.RedirectUrl = ""
	v.BindSiteUserId = 0
	v.EntId = 0
	v.OperateTime = 0
	poolNrtEaLoginDto.Put(v)
}
