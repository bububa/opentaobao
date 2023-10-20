package alihealthoutflow

import (
	"sync"
)

// DoctorStatusDto 结构体
type DoctorStatusDto struct {
	// 审核拒绝原因/证书信息
	Note string `json:"note,omitempty" xml:"note,omitempty"`
	// 用户同步状态 0：身份审核通过 1：证书签发 2：设置签章 3：用户注销 4：申请拒绝 5：用户停用 6：修改手机号 7：用户启用 9：用户删除
	Process string `json:"process,omitempty" xml:"process,omitempty"`
	// 第三方厂商标识（必填）
	ClientId string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// 医网信医师唯一标识
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 整体参数签名值
	Sign string `json:"sign,omitempty" xml:"sign,omitempty"`
	// 仅当process为2时，会回调签章图片base64
	Stamp string `json:"stamp,omitempty" xml:"stamp,omitempty"`
	// 用户手机号
	PhoneNum string `json:"phone_num,omitempty" xml:"phone_num,omitempty"`
	// 操作时间，yyyy-MM-dd HH:mm:ss
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 仅当process为2时，会回调签章状态 10：待审核 11：签章审核通过
	StampStatus string `json:"stamp_status,omitempty" xml:"stamp_status,omitempty"`
	// 模板ID（必填）
	TemplateId string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 订阅服务ID（必填）
	ServiceId int64 `json:"service_id,omitempty" xml:"service_id,omitempty"`
}

var poolDoctorStatusDto = sync.Pool{
	New: func() any {
		return new(DoctorStatusDto)
	},
}

// GetDoctorStatusDto() 从对象池中获取DoctorStatusDto
func GetDoctorStatusDto() *DoctorStatusDto {
	return poolDoctorStatusDto.Get().(*DoctorStatusDto)
}

// ReleaseDoctorStatusDto 释放DoctorStatusDto
func ReleaseDoctorStatusDto(v *DoctorStatusDto) {
	v.Note = ""
	v.Process = ""
	v.ClientId = ""
	v.OpenId = ""
	v.Sign = ""
	v.Stamp = ""
	v.PhoneNum = ""
	v.Time = ""
	v.StampStatus = ""
	v.TemplateId = ""
	v.ServiceId = 0
	poolDoctorStatusDto.Put(v)
}
