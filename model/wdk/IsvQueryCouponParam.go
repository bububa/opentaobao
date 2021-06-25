package wdk

// IsvQueryCouponParam 
type IsvQueryCouponParam struct {

    // umpId列表，最多支持一次批量查询20个
    UmpIdList   []Number `json:"ump_id_list,omitempty"`

}
