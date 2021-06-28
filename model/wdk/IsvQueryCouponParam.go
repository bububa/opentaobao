package wdk

// IsvQueryCouponParam 
/* model for simplify = false
type IsvQueryCouponParam struct {

    // umpId列表，最多支持一次批量查询20个
    
    UmpIdList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"ump_id_list,omitempty"`
    

}
*/

// IsvQueryCouponParam 
type IsvQueryCouponParam struct {

    // umpId列表，最多支持一次批量查询20个
    UmpIdList   []int64 `json:"ump_id_list,omitempty"`

}
