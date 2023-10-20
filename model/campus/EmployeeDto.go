package campus

import (
	"sync"
)

// EmployeeDto 结构体
type EmployeeDto struct {
	// 生日
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 性别
	Sex string `json:"sex,omitempty" xml:"sex,omitempty"`
	// 钉钉
	Dingding string `json:"dingding,omitempty" xml:"dingding,omitempty"`
	// 旺旺
	Wangwang string `json:"wangwang,omitempty" xml:"wangwang,omitempty"`
	// 英文名称
	EnName string `json:"en_name,omitempty" xml:"en_name,omitempty"`
	// 微信
	Weixin string `json:"weixin,omitempty" xml:"weixin,omitempty"`
	// 别名
	NickName string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// 手机号备注，一般用于存放第二个手机号
	MobileComment string `json:"mobile_comment,omitempty" xml:"mobile_comment,omitempty"`
	// 钉钉UserId
	DingUserId string `json:"ding_user_id,omitempty" xml:"ding_user_id,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 头像Url
	AvatarUrl string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
	// qq
	Qq string `json:"qq,omitempty" xml:"qq,omitempty"`
	// 公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 头像
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 工号
	WorkNo string `json:"work_no,omitempty" xml:"work_no,omitempty"`
	// 支付宝
	Alipay string `json:"alipay,omitempty" xml:"alipay,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// corp_user_id
	CorpUserId string `json:"corp_user_id,omitempty" xml:"corp_user_id,omitempty"`
	// 头像Url
	AvatarPreViewUrl string `json:"avatar_pre_view_url,omitempty" xml:"avatar_pre_view_url,omitempty"`
	// 工作状态
	WorkStatus string `json:"work_status,omitempty" xml:"work_status,omitempty"`
	// 座机号码
	Telephone string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 手机
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 用户id
	AccountId int64 `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 公司ID
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
	// 用户账号id
	PassportAccountId int64 `json:"passport_account_id,omitempty" xml:"passport_account_id,omitempty"`
}

var poolEmployeeDto = sync.Pool{
	New: func() any {
		return new(EmployeeDto)
	},
}

// GetEmployeeDto() 从对象池中获取EmployeeDto
func GetEmployeeDto() *EmployeeDto {
	return poolEmployeeDto.Get().(*EmployeeDto)
}

// ReleaseEmployeeDto 释放EmployeeDto
func ReleaseEmployeeDto(v *EmployeeDto) {
	v.Birthday = ""
	v.Sex = ""
	v.Dingding = ""
	v.Wangwang = ""
	v.EnName = ""
	v.Weixin = ""
	v.NickName = ""
	v.MobileComment = ""
	v.DingUserId = ""
	v.Name = ""
	v.AvatarUrl = ""
	v.Qq = ""
	v.CompanyName = ""
	v.Avatar = ""
	v.WorkNo = ""
	v.Alipay = ""
	v.Email = ""
	v.CorpUserId = ""
	v.AvatarPreViewUrl = ""
	v.WorkStatus = ""
	v.Telephone = ""
	v.Mobile = ""
	v.AccountId = 0
	v.Id = 0
	v.CompanyId = 0
	v.PassportAccountId = 0
	poolEmployeeDto.Put(v)
}
