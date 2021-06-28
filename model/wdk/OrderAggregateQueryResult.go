package wdk

// OrderAggregateQueryResult 
/* model for simplify = false
type OrderAggregateQueryResult struct {

    // 接口返回码. 如果返回 HM05038888888006 需重试(数据查询失败，请重试，注意限定重试次数)
    
    ReturnCode   string `json:"return_code,omitempty"`
    

    // 接口返回码描述
    
    ReturnMsg   string `json:"return_msg,omitempty"`
    

    // 240000310869037498
    
    BizIdList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"biz_id_list,omitempty"`
    

    // 171321816752430897
    
    TbBizIdList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"tb_biz_id_list,omitempty"`
    

    // totalNum
    
    TotalNum   int64 `json:"total_num,omitempty"`
    

    // 商品总金额（不含优惠）
    
    OriginalAmt   int64 `json:"original_amt,omitempty"`
    

    // 总优惠金额
    
    DiscountAmt   int64 `json:"discount_amt,omitempty"`
    

    // 下一页序号
    
    NextIndex   int64 `json:"next_index,omitempty"`
    

}
*/

// OrderAggregateQueryResult 
type OrderAggregateQueryResult struct {

    // 接口返回码. 如果返回 HM05038888888006 需重试(数据查询失败，请重试，注意限定重试次数)
    ReturnCode   string `json:"return_code,omitempty"`

    // 接口返回码描述
    ReturnMsg   string `json:"return_msg,omitempty"`

    // 240000310869037498
    BizIdList   []int64 `json:"biz_id_list,omitempty"`

    // 171321816752430897
    TbBizIdList   []int64 `json:"tb_biz_id_list,omitempty"`

    // totalNum
    TotalNum   int64 `json:"total_num,omitempty"`

    // 商品总金额（不含优惠）
    OriginalAmt   int64 `json:"original_amt,omitempty"`

    // 总优惠金额
    DiscountAmt   int64 `json:"discount_amt,omitempty"`

    // 下一页序号
    NextIndex   int64 `json:"next_index,omitempty"`

}
