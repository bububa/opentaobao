package btrip

import (
	"sync"
)

// OpenCostCenterQueryRs 结构体
type OpenCostCenterQueryRs struct {
	// 绑定人员信息
	EntityList []OpenOrgEntityDo `json:"entity_list,omitempty" xml:"entity_list>open_org_entity_do,omitempty"`
	// 企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 成本中心名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 成本中心编号
	Number string `json:"number,omitempty" xml:"number,omitempty"`
	// 第三方成本中心id
	ThirdpartId string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	// 绑定支付宝账号
	AlipayNo string `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
	// 商旅成本中心id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 适用范围: 1全员，2部分员工
	Scope int64 `json:"scope,omitempty" xml:"scope,omitempty"`
}

var poolOpenCostCenterQueryRs = sync.Pool{
	New: func() any {
		return new(OpenCostCenterQueryRs)
	},
}

// GetOpenCostCenterQueryRs() 从对象池中获取OpenCostCenterQueryRs
func GetOpenCostCenterQueryRs() *OpenCostCenterQueryRs {
	return poolOpenCostCenterQueryRs.Get().(*OpenCostCenterQueryRs)
}

// ReleaseOpenCostCenterQueryRs 释放OpenCostCenterQueryRs
func ReleaseOpenCostCenterQueryRs(v *OpenCostCenterQueryRs) {
	v.EntityList = v.EntityList[:0]
	v.CorpId = ""
	v.Title = ""
	v.Number = ""
	v.ThirdpartId = ""
	v.AlipayNo = ""
	v.Id = 0
	v.Scope = 0
	poolOpenCostCenterQueryRs.Put(v)
}
