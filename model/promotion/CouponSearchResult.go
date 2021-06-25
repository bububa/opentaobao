package promotion

// CouponSearchResult 
type CouponSearchResult struct {

    // 符合条件总数量，用于分页等判断
    TotalCount   int64 `json:"total_count,omitempty"`

    // 排查调用id
    TraceId   string `json:"trace_id,omitempty"`

    // 优惠券详情列表
    SellerCouponDetails   []AllsparkSellerCouponDetail `json:"seller_coupon_details,omitempty"`

}
