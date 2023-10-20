package nrt

import (
	"sync"
)

// NrtMemberDto 结构体
type NrtMemberDto struct {
	// 手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 业务身份
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 验证码
	SmsCode string `json:"sms_code,omitempty" xml:"sms_code,omitempty"`
	// 外部会员ID
	OutMemberId string `json:"out_member_id,omitempty" xml:"out_member_id,omitempty"`
	// 淘宝ID
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 操作类型
	OpType string `json:"op_type,omitempty" xml:"op_type,omitempty"`
	// 会员等级
	LevelName string `json:"level_name,omitempty" xml:"level_name,omitempty"`
	// 会员名称
	RealName string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// 幂等ID
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 会员地址
	Addr string `json:"addr,omitempty" xml:"addr,omitempty"`
	// 状态 1:正常 0:冻结
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 淘宝nick
	TaoNick string `json:"tao_nick,omitempty" xml:"tao_nick,omitempty"`
	// 会员类型（1，喵零；2会员通）
	MemberType int64 `json:"member_type,omitempty" xml:"member_type,omitempty"`
	// 同步时间
	GmtModified int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolNrtMemberDto = sync.Pool{
	New: func() any {
		return new(NrtMemberDto)
	},
}

// GetNrtMemberDto() 从对象池中获取NrtMemberDto
func GetNrtMemberDto() *NrtMemberDto {
	return poolNrtMemberDto.Get().(*NrtMemberDto)
}

// ReleaseNrtMemberDto 释放NrtMemberDto
func ReleaseNrtMemberDto(v *NrtMemberDto) {
	v.Phone = ""
	v.BizCode = ""
	v.SmsCode = ""
	v.OutMemberId = ""
	v.OpenId = ""
	v.OpType = ""
	v.LevelName = ""
	v.RealName = ""
	v.OutId = ""
	v.Addr = ""
	v.Status = ""
	v.TaoNick = ""
	v.MemberType = 0
	v.GmtModified = 0
	v.Version = 0
	poolNrtMemberDto.Put(v)
}
