package alihealthoutflow

import (
	"sync"
)

// CaSignStatusSaveRequest 结构体
type CaSignStatusSaveRequest struct {
	// 子机构id
	ThirdDepartId string `json:"third_depart_id,omitempty" xml:"third_depart_id,omitempty"`
	// 第三方下属厂商id(最大64位)
	ClientId string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// 证书有效期结束时间
	CertEndDate string `json:"cert_end_date,omitempty" xml:"cert_end_date,omitempty"`
	// 签名值
	Sign string `json:"sign,omitempty" xml:"sign,omitempty"`
	// 公章base64编码格式(仅在制作签章时回调)
	Stamp string `json:"stamp,omitempty" xml:"stamp,omitempty"`
	// 子机构在医网信id(子机构审核制章的值)
	DepartId string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// 描述(拒绝有值)
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 证书序列号
	CertNum string `json:"cert_num,omitempty" xml:"cert_num,omitempty"`
	// 证书有效期开始时间
	CertStartDate string `json:"cert_start_date,omitempty" xml:"cert_start_date,omitempty"`
	// 状态 1:通过 2:停用 3:拒绝
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolCaSignStatusSaveRequest = sync.Pool{
	New: func() any {
		return new(CaSignStatusSaveRequest)
	},
}

// GetCaSignStatusSaveRequest() 从对象池中获取CaSignStatusSaveRequest
func GetCaSignStatusSaveRequest() *CaSignStatusSaveRequest {
	return poolCaSignStatusSaveRequest.Get().(*CaSignStatusSaveRequest)
}

// ReleaseCaSignStatusSaveRequest 释放CaSignStatusSaveRequest
func ReleaseCaSignStatusSaveRequest(v *CaSignStatusSaveRequest) {
	v.ThirdDepartId = ""
	v.ClientId = ""
	v.CertEndDate = ""
	v.Sign = ""
	v.Stamp = ""
	v.DepartId = ""
	v.Description = ""
	v.CertNum = ""
	v.CertStartDate = ""
	v.Status = ""
	poolCaSignStatusSaveRequest.Put(v)
}
