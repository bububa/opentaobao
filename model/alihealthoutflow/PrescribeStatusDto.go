package alihealthoutflow

import (
	"sync"
)

// PrescribeStatusDto 结构体
type PrescribeStatusDto struct {
	// 拒绝签名原因（仅在拒绝签名状态时回调）
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 第三方厂商标识
	ClientId string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// pdf文件流（base64编码，（仅在pdf签名成功状态时回调））
	SignedPdfBase64 string `json:"signed_pdf_base64,omitempty" xml:"signed_pdf_base64,omitempty"`
	// 作废时间（仅在作废状态时回调），yyyy-MM-dd HH:mm:ss
	DeleteTime string `json:"delete_time,omitempty" xml:"delete_time,omitempty"`
	// 处方哈希的 P7签名（仅在签名成功时回调）
	SignedData string `json:"signed_data,omitempty" xml:"signed_data,omitempty"`
	// 签名医师openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// Body签名值
	Sign string `json:"sign,omitempty" xml:"sign,omitempty"`
	// 模板Id
	TemplateId string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 处方ID
	UrId string `json:"ur_id,omitempty" xml:"ur_id,omitempty"`
	// 医网信处方唯一标识
	UniqueId string `json:"unique_id,omitempty" xml:"unique_id,omitempty"`
	// 订阅服务Id
	ServiceId int64 `json:"service_id,omitempty" xml:"service_id,omitempty"`
	// 签名订单状态 2已签名 6拒绝签名 7签名订单已过期删除 9已签名订单作废
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolPrescribeStatusDto = sync.Pool{
	New: func() any {
		return new(PrescribeStatusDto)
	},
}

// GetPrescribeStatusDto() 从对象池中获取PrescribeStatusDto
func GetPrescribeStatusDto() *PrescribeStatusDto {
	return poolPrescribeStatusDto.Get().(*PrescribeStatusDto)
}

// ReleasePrescribeStatusDto 释放PrescribeStatusDto
func ReleasePrescribeStatusDto(v *PrescribeStatusDto) {
	v.Reason = ""
	v.ClientId = ""
	v.SignedPdfBase64 = ""
	v.DeleteTime = ""
	v.SignedData = ""
	v.OpenId = ""
	v.Sign = ""
	v.TemplateId = ""
	v.UrId = ""
	v.UniqueId = ""
	v.ServiceId = 0
	v.Status = 0
	poolPrescribeStatusDto.Put(v)
}
