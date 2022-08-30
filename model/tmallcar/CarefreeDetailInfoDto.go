package tmallcar

// CarefreeDetailInfoDto 结构体
type CarefreeDetailInfoDto struct {
	// 商品标题
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 商品图片
	ItemPic string `json:"item_pic,omitempty" xml:"item_pic,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 车系名称
	SeriesName string `json:"series_name,omitempty" xml:"series_name,omitempty"`
	// 核销门店名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 业务单号
	BizId int64 `json:"biz_id,omitempty" xml:"biz_id,omitempty"`
	// 支付状态
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 未上传发票服务失效时间（毫秒）
	ServiceOffTime int64 `json:"service_off_time,omitempty" xml:"service_off_time,omitempty"`
	// 是否发生退款
	HasRefunded int64 `json:"has_refunded,omitempty" xml:"has_refunded,omitempty"`
	// 品牌id
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 车系id
	SeriesId int64 `json:"series_id,omitempty" xml:"series_id,omitempty"`
	// 核销门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 订单类型（1 - 一口价电子凭证订单 | 2 - 分阶段订单）
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 定金支付时间
	EarnestPaidTime int64 `json:"earnest_paid_time,omitempty" xml:"earnest_paid_time,omitempty"`
	// 订单关闭时间
	OrderCloseTime int64 `json:"order_close_time,omitempty" xml:"order_close_time,omitempty"`
	// orderCancelTime
	OrderCancelTime int64 `json:"order_cancel_time,omitempty" xml:"order_cancel_time,omitempty"`
	// price
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// orderSuccessTime
	OrderSuccessTime int64 `json:"order_success_time,omitempty" xml:"order_success_time,omitempty"`
	// 是否免申请
	FreeApply bool `json:"free_apply,omitempty" xml:"free_apply,omitempty"`
}
