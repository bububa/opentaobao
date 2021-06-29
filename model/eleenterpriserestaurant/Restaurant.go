package eleenterpriserestaurant

// Restaurant 
type Restaurant struct {
    // 距离
    Distance   string `json:"distance,omitempty" xml:"distance,omitempty"`
    // 评价
    Rating   string `json:"rating,omitempty" xml:"rating,omitempty"`
    // 唯一店铺标识
    OnlyRestaurantCode   string `json:"only_restaurant_code,omitempty" xml:"only_restaurant_code,omitempty"`
    // 促销信息
    PromotionInfo   string `json:"promotion_info,omitempty" xml:"promotion_info,omitempty"`
    // 最近一个月订单量
    RecentOrderNum   int64 `json:"recent_order_num,omitempty" xml:"recent_order_num,omitempty"`
    // 配送费
    AgentFee   string `json:"agent_fee,omitempty" xml:"agent_fee,omitempty"`
    // 2周内的平均送餐时间
    DeliverSpent   int64 `json:"deliver_spent,omitempty" xml:"deliver_spent,omitempty"`
    // 是否正在营业
    IsOpen   int64 `json:"is_open,omitempty" xml:"is_open,omitempty"`
    // 是否蜂鸟专送餐厅
    IsDistRst   int64 `json:"is_dist_rst,omitempty" xml:"is_dist_rst,omitempty"`
    // 餐厅名称
    RestaurantName   string `json:"restaurant_name,omitempty" xml:"restaurant_name,omitempty"`
    // 活动详情
    Activities   []Activities `json:"activities,omitempty" xml:"activities>activities,omitempty"`
    // 餐厅Logo地址
    ImageUrl   string `json:"image_url,omitempty" xml:"image_url,omitempty"`
    // 起送价
    DeliverAmount   string `json:"deliver_amount,omitempty" xml:"deliver_amount,omitempty"`
    // 店铺id
    ErestaurantId   string `json:"erestaurant_id,omitempty" xml:"erestaurant_id,omitempty"`
    // 是否支持开发票(是为1，否为0)
    Invoice   int64 `json:"invoice,omitempty" xml:"invoice,omitempty"`
    // 是否品牌馆餐厅
    IsPremium   int64 `json:"is_premium,omitempty" xml:"is_premium,omitempty"`
    // 人均值
    AverageCost   string `json:"average_cost,omitempty" xml:"average_cost,omitempty"`
    // 是否支持食安保：0，否；1，是
    IsInsurance   int64 `json:"is_insurance,omitempty" xml:"is_insurance,omitempty"`
    // 是否新开业
    IsNew   int64 `json:"is_new,omitempty" xml:"is_new,omitempty"`
    // 水印序列号
    SerialNumber   string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
    // 餐厅整体营业状态：1 餐厅营业中，2餐厅关闭，3 餐厅网路不稳定，4 餐厅休息中，5 直接说预定，6 只接受电话预定，7 餐厅休假中
    TotalStatus   string `json:"total_status,omitempty" xml:"total_status,omitempty"`
    // 是否支持预定. 0 不支持预定, 1 支持预定
    IsBookable   int64 `json:"is_bookable,omitempty" xml:"is_bookable,omitempty"`
}
