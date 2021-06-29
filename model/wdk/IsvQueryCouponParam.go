package wdk

// IsvQueryCouponParam 
type IsvQueryCouponParam struct {
    // umpId列表，最多支持一次批量查询20个
    UmpIdList   []int64 `json:"ump_id_list,omitempty" xml:"ump_id_list>int64,omitempty"`
}
