package promotion

import (
	"sync"
)

// CouponTemplateOperateResponse 结构体
type CouponTemplateOperateResponse struct {
	// 创建结果
	FailElements []FailElement `json:"fail_elements,omitempty" xml:"fail_elements>fail_element,omitempty"`
	// 券圈品设置
	PromActSkuList []PromActSku `json:"prom_act_sku_list,omitempty" xml:"prom_act_sku_list>prom_act_sku,omitempty"`
	// 匿名码code
	MaCode string `json:"ma_code,omitempty" xml:"ma_code,omitempty"`
	// 券实例id
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 模板表主键
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// ump模板ID
	SourceId int64 `json:"source_id,omitempty" xml:"source_id,omitempty"`
	// 分页信息
	PageInfo *PageInfo `json:"page_info,omitempty" xml:"page_info,omitempty"`
	// 券模版
	CouponTemplate *CouponTemplate `json:"coupon_template,omitempty" xml:"coupon_template,omitempty"`
}

var poolCouponTemplateOperateResponse = sync.Pool{
	New: func() any {
		return new(CouponTemplateOperateResponse)
	},
}

// GetCouponTemplateOperateResponse() 从对象池中获取CouponTemplateOperateResponse
func GetCouponTemplateOperateResponse() *CouponTemplateOperateResponse {
	return poolCouponTemplateOperateResponse.Get().(*CouponTemplateOperateResponse)
}

// ReleaseCouponTemplateOperateResponse 释放CouponTemplateOperateResponse
func ReleaseCouponTemplateOperateResponse(v *CouponTemplateOperateResponse) {
	v.FailElements = v.FailElements[:0]
	v.PromActSkuList = v.PromActSkuList[:0]
	v.MaCode = ""
	v.VoucherId = ""
	v.Id = 0
	v.SourceId = 0
	v.PageInfo = nil
	v.CouponTemplate = nil
	poolCouponTemplateOperateResponse.Put(v)
}
