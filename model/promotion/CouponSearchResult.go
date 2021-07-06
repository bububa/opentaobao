package promotion

// CouponSearchResult 结构体
type CouponSearchResult struct {
	// 优惠券详情列表
	SellerCouponDetails []AllsparkSellerCouponDetail `json:"seller_coupon_details,omitempty" xml:"seller_coupon_details>allspark_seller_coupon_detail,omitempty"`
	// 排查调用id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 符合条件总数量，用于分页等判断
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
