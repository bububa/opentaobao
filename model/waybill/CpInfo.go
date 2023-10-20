package waybill

import (
	"sync"
)

// CpInfo 结构体
type CpInfo struct {
	// 云打印模板
	CloudTemplateId string `json:"cloud_template_id,omitempty" xml:"cloud_template_id,omitempty"`
	// 快递公司
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 地址信息
	Address *Address `json:"address,omitempty" xml:"address,omitempty"`
	// 状态: 0-禁用, 1-启用
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolCpInfo = sync.Pool{
	New: func() any {
		return new(CpInfo)
	},
}

// GetCpInfo() 从对象池中获取CpInfo
func GetCpInfo() *CpInfo {
	return poolCpInfo.Get().(*CpInfo)
}

// ReleaseCpInfo 释放CpInfo
func ReleaseCpInfo(v *CpInfo) {
	v.CloudTemplateId = ""
	v.CpCode = ""
	v.Address = nil
	v.Status = 0
	poolCpInfo.Put(v)
}
