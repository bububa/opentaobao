package idle

import (
	"sync"
)

// InspectionReport 结构体
type InspectionReport struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 鉴定报告
	Report string `json:"report,omitempty" xml:"report,omitempty"`
	// 错误描述
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 成色
	Degree string `json:"degree,omitempty" xml:"degree,omitempty"`
	// 估价id,与order_id任选其一，orderId优先
	QuoteId string `json:"quote_id,omitempty" xml:"quote_id,omitempty"`
	// 对此商品的质检描述
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// *      * Json串:desc存储特别说明的文字描述,images存储上传的图片url;      * 形如:      * {      *   &#34;desc&#34;:&#34;desc&#34;,      *   &#34;images&#34;:[      *      &#34;imageUrl1&#34;,      *      &#34;imageUrl2&#34;      *   ]      * };
	Explanation string `json:"explanation,omitempty" xml:"explanation,omitempty"`
	// 设备imei号
	Imei string `json:"imei,omitempty" xml:"imei,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 订单，orderId优先，与quote_id任选其一
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 手淘商家的淘宝账号id，通过 spuId 和 recycleSupplierId 定位到唯一的回收模板
	RecycleSupplierId int64 `json:"recycle_supplier_id,omitempty" xml:"recycle_supplier_id,omitempty"`
	// 质检类型（1第一次直接;2二次复检）
	AppraiseType int64 `json:"appraise_type,omitempty" xml:"appraise_type,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolInspectionReport = sync.Pool{
	New: func() any {
		return new(InspectionReport)
	},
}

// GetInspectionReport() 从对象池中获取InspectionReport
func GetInspectionReport() *InspectionReport {
	return poolInspectionReport.Get().(*InspectionReport)
}

// ReleaseInspectionReport 释放InspectionReport
func ReleaseInspectionReport(v *InspectionReport) {
	v.ErrMsg = ""
	v.Report = ""
	v.ErrCode = ""
	v.Degree = ""
	v.QuoteId = ""
	v.Summary = ""
	v.Explanation = ""
	v.Imei = ""
	v.Price = 0
	v.OrderId = 0
	v.RecycleSupplierId = 0
	v.AppraiseType = 0
	v.Success = false
	poolInspectionReport.Put(v)
}
