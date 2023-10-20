package alicom

import (
	"sync"
)

// PreAuthorizeModel 结构体
type PreAuthorizeModel struct {
	// 扩展字段(json)
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 淘宝订单号
	TbOrderNo string `json:"tb_order_no,omitempty" xml:"tb_order_no,omitempty"`
	// 外部流水号
	OutBizOrderNo string `json:"out_biz_order_no,omitempty" xml:"out_biz_order_no,omitempty"`
	// 授权水号
	FundAuthNo string `json:"fund_auth_no,omitempty" xml:"fund_auth_no,omitempty"`
	// 淘宝订单号
	Tid string `json:"tid,omitempty" xml:"tid,omitempty"`
	// 0:业务办理成功，2:业务办理失败,全额解冻
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolPreAuthorizeModel = sync.Pool{
	New: func() any {
		return new(PreAuthorizeModel)
	},
}

// GetPreAuthorizeModel() 从对象池中获取PreAuthorizeModel
func GetPreAuthorizeModel() *PreAuthorizeModel {
	return poolPreAuthorizeModel.Get().(*PreAuthorizeModel)
}

// ReleasePreAuthorizeModel 释放PreAuthorizeModel
func ReleasePreAuthorizeModel(v *PreAuthorizeModel) {
	v.Ext = ""
	v.TbOrderNo = ""
	v.OutBizOrderNo = ""
	v.FundAuthNo = ""
	v.Tid = ""
	v.Status = 0
	poolPreAuthorizeModel.Put(v)
}
