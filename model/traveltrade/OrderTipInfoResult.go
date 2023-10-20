package traveltrade

import (
	"sync"
)

// OrderTipInfoResult 结构体
type OrderTipInfoResult struct {
	// 查看模版链接
	ViewTemplateUrl string `json:"view_template_url,omitempty" xml:"view_template_url,omitempty"`
	// 模版详情
	TemplateInfo *OrderTemplateInfo `json:"template_info,omitempty" xml:"template_info,omitempty"`
	// 是否是最新的模版，如果订单没有标注模版，则默认取最新的模版数据；如果订单有标注模版，则取对应的模版信息；
	IsNew bool `json:"is_new,omitempty" xml:"is_new,omitempty"`
}

var poolOrderTipInfoResult = sync.Pool{
	New: func() any {
		return new(OrderTipInfoResult)
	},
}

// GetOrderTipInfoResult() 从对象池中获取OrderTipInfoResult
func GetOrderTipInfoResult() *OrderTipInfoResult {
	return poolOrderTipInfoResult.Get().(*OrderTipInfoResult)
}

// ReleaseOrderTipInfoResult 释放OrderTipInfoResult
func ReleaseOrderTipInfoResult(v *OrderTipInfoResult) {
	v.ViewTemplateUrl = ""
	v.TemplateInfo = nil
	v.IsNew = false
	poolOrderTipInfoResult.Put(v)
}
