package qt

// ServiceSubscribe 
/* model for simplify = false
type ServiceSubscribe struct {

    // 服务收费项code
    
    ServiceItemCode   string `json:"service_item_code,omitempty"`
    

    // 订购者昵称
    
    Nick   string `json:"nick,omitempty"`
    

    // 订购总数
    
    AllNum   int64 `json:"all_num,omitempty"`
    

    // 已经使用的数量
    
    UsedNum   int64 `json:"used_num,omitempty"`
    

    // 过期时间
    
    GmtExpiry   string `json:"gmt_expiry,omitempty"`
    

    // 将要被使用的那条质检订单的价格
    
    FuturePrice   float64 `json:"future_price,omitempty"`
    

    // 该用户该收费项目下面的所有的订购记录详情
    
    UsageDetailList  struct {
        QualityUsageDetail  []QualityUsageDetail `json:"quality_usage_detail,omitempty"`
    } `json:"usage_detail_list,omitempty"`
    

    // 将要被消耗的质检订单ID
    
    FutureSubId   int64 `json:"future_sub_id,omitempty"`
    

    // 可用数量
    
    AvaliableNum   int64 `json:"avaliable_num,omitempty"`
    

}
*/

// ServiceSubscribe 
type ServiceSubscribe struct {

    // 服务收费项code
    ServiceItemCode   string `json:"service_item_code,omitempty"`

    // 订购者昵称
    Nick   string `json:"nick,omitempty"`

    // 订购总数
    AllNum   int64 `json:"all_num,omitempty"`

    // 已经使用的数量
    UsedNum   int64 `json:"used_num,omitempty"`

    // 过期时间
    GmtExpiry   string `json:"gmt_expiry,omitempty"`

    // 将要被使用的那条质检订单的价格
    FuturePrice   float64 `json:"future_price,omitempty"`

    // 该用户该收费项目下面的所有的订购记录详情
    UsageDetailList   []QualityUsageDetail `json:"usage_detail_list,omitempty"`

    // 将要被消耗的质检订单ID
    FutureSubId   int64 `json:"future_sub_id,omitempty"`

    // 可用数量
    AvaliableNum   int64 `json:"avaliable_num,omitempty"`

}
