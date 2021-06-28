package alicom

// SupplierTopQueryModel 
/* model for simplify = false
type SupplierTopQueryModel struct {

    // 业务类型:7-合约及号卡分销
    
    BizType   string `json:"biz_type,omitempty"`
    

    // 分销商名称
    
    DistributorName   string `json:"distributor_name,omitempty"`
    

    // 订单结束时间
    
    EndTime   string `json:"end_time,omitempty"`
    

    // 淘宝交易订单号
    
    OrderNo   string `json:"order_no,omitempty"`
    

    // 订单状态列表:1-未订购,2-订购中,3-订购中,4-订购失败,5-订购成功,6-订购取消
    
    OrderStatusList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"order_status_list,omitempty"`
    

    // 当前页
    
    PageNum   int64 `json:"page_num,omitempty"`
    

    // 分页数量
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 用户账号
    
    PhoneNo   string `json:"phone_no,omitempty"`
    

    // 订单开始时间
    
    StartTime   string `json:"start_time,omitempty"`
    

}
*/

// SupplierTopQueryModel 
type SupplierTopQueryModel struct {

    // 业务类型:7-合约及号卡分销
    BizType   string `json:"biz_type,omitempty"`

    // 分销商名称
    DistributorName   string `json:"distributor_name,omitempty"`

    // 订单结束时间
    EndTime   string `json:"end_time,omitempty"`

    // 淘宝交易订单号
    OrderNo   string `json:"order_no,omitempty"`

    // 订单状态列表:1-未订购,2-订购中,3-订购中,4-订购失败,5-订购成功,6-订购取消
    OrderStatusList   []int64 `json:"order_status_list,omitempty"`

    // 当前页
    PageNum   int64 `json:"page_num,omitempty"`

    // 分页数量
    PageSize   int64 `json:"page_size,omitempty"`

    // 用户账号
    PhoneNo   string `json:"phone_no,omitempty"`

    // 订单开始时间
    StartTime   string `json:"start_time,omitempty"`

}
