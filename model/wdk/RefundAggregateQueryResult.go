package wdk

// RefundAggregateQueryResult 
/* model for simplify = false
type RefundAggregateQueryResult struct {

    // 接口返回码. 如果返回 HM05038888888006 需重试(数据查询失败，请重试，注意限定重试次数)
    
    ReturnCode   string `json:"return_code,omitempty"`
    

    // 接口返回码描述
    
    ReturnMsg   string `json:"return_msg,omitempty"`
    

    // 退款单id列表
    
    RefundIdList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"refund_id_list,omitempty"`
    

    // 淘鲜达子订单id列表
    
    BizIdList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"biz_id_list,omitempty"`
    

    // 淘宝子订单id列表
    
    TbBizIdList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"tb_biz_id_list,omitempty"`
    

    // 结果数量
    
    TotalNum   int64 `json:"total_num,omitempty"`
    

    // 下一页序号
    
    NextIndex   int64 `json:"next_index,omitempty"`
    

}
*/

// RefundAggregateQueryResult 
type RefundAggregateQueryResult struct {

    // 接口返回码. 如果返回 HM05038888888006 需重试(数据查询失败，请重试，注意限定重试次数)
    ReturnCode   string `json:"return_code,omitempty"`

    // 接口返回码描述
    ReturnMsg   string `json:"return_msg,omitempty"`

    // 退款单id列表
    RefundIdList   []int64 `json:"refund_id_list,omitempty"`

    // 淘鲜达子订单id列表
    BizIdList   []int64 `json:"biz_id_list,omitempty"`

    // 淘宝子订单id列表
    TbBizIdList   []int64 `json:"tb_biz_id_list,omitempty"`

    // 结果数量
    TotalNum   int64 `json:"total_num,omitempty"`

    // 下一页序号
    NextIndex   int64 `json:"next_index,omitempty"`

}
