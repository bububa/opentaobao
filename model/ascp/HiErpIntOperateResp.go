package ascp

import (
	"sync"
)

// HiErpIntOperateResp 结构体
type HiErpIntOperateResp struct {
	// 履约单号
	ScpCode string `json:"scp_code,omitempty" xml:"scp_code,omitempty"`
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 拦截结果文案描述
	TmsCutResultText string `json:"tms_cut_result_text,omitempty" xml:"tms_cut_result_text,omitempty"`
	// 外部订单号
	OutOrderCode string `json:"out_order_code,omitempty" xml:"out_order_code,omitempty"`
}

var poolHiErpIntOperateResp = sync.Pool{
	New: func() any {
		return new(HiErpIntOperateResp)
	},
}

// GetHiErpIntOperateResp() 从对象池中获取HiErpIntOperateResp
func GetHiErpIntOperateResp() *HiErpIntOperateResp {
	return poolHiErpIntOperateResp.Get().(*HiErpIntOperateResp)
}

// ReleaseHiErpIntOperateResp 释放HiErpIntOperateResp
func ReleaseHiErpIntOperateResp(v *HiErpIntOperateResp) {
	v.ScpCode = ""
	v.MailNo = ""
	v.TmsCutResultText = ""
	v.OutOrderCode = ""
	poolHiErpIntOperateResp.Put(v)
}
