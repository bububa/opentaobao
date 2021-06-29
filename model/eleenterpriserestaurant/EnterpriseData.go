package eleenterpriserestaurant

// EnterpriseData 
type EnterpriseData struct {

    // 类目名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 1-餐饮，2-新零售，目前只返回餐饮类目
    
    Channel   int64 `json:"channel,omitempty" xml:"channel,omitempty"`
    

    // 类目id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 父级类目id，目前只返回一级类目id，都为0。预留字段，用于今后扩展
    
    ParentId   int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
    

    // 餐厅维度
    
    Latitude   string `json:"latitude,omitempty" xml:"latitude,omitempty"`
    

    // 是否可配送
    
    IsDeliverable   bool `json:"is_deliverable,omitempty" xml:"is_deliverable,omitempty"`
    

    // 餐厅精度
    
    Longitude   string `json:"longitude,omitempty" xml:"longitude,omitempty"`
    

    // 距离（米）
    
    Distance   string `json:"distance,omitempty" xml:"distance,omitempty"`
    

    // 店铺综合评分
    
    Rating   string `json:"rating,omitempty" xml:"rating,omitempty"`
    

    // 餐厅唯一码
    
    OnlyRestaurantCode   string `json:"only_restaurant_code,omitempty" xml:"only_restaurant_code,omitempty"`
    

    // 促销信息或者商家公告
    
    PromotionInfo   string `json:"promotion_info,omitempty" xml:"promotion_info,omitempty"`
    

    // 最近一个月订单量
    
    RecentOrderNum   int64 `json:"recent_order_num,omitempty" xml:"recent_order_num,omitempty"`
    

    // 配送费
    
    AgentFee   string `json:"agent_fee,omitempty" xml:"agent_fee,omitempty"`
    

    // 2周内的平均送餐时间
    
    DeliverSpent   int64 `json:"deliver_spent,omitempty" xml:"deliver_spent,omitempty"`
    

    // 是否正在营业, 在营业的前提下，再通过total_status字段判断商家的具体营业状态 餐厅整体营业状态
    
    IsOpen   int64 `json:"is_open,omitempty" xml:"is_open,omitempty"`
    

    // 是否蜂鸟专送餐厅
    
    IsDistRst   int64 `json:"is_dist_rst,omitempty" xml:"is_dist_rst,omitempty"`
    

    // 餐厅名称
    
    RestaurantName   string `json:"restaurant_name,omitempty" xml:"restaurant_name,omitempty"`
    

    // 活动, 参考餐厅活动
    
    Activities   []Activitie `json:"activities,omitempty" xml:"activities,omitempty"`
    

    // 餐厅Logo地址
    
    ImageUrl   string `json:"image_url,omitempty" xml:"image_url,omitempty"`
    

    // 起送价
    
    DeliverAmount   string `json:"deliver_amount,omitempty" xml:"deliver_amount,omitempty"`
    

    // 餐厅ID
    
    ErestaurantId   string `json:"erestaurant_id,omitempty" xml:"erestaurant_id,omitempty"`
    

    // 是否支持开发票: 1, 是; 2, 否.
    
    Invoice   int64 `json:"invoice,omitempty" xml:"invoice,omitempty"`
    

    // 序列号（无业务含义）
    
    SerialNumber   string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
    

    // 是否支持预定. 0 不支持预定, 1 支持预定
    
    IsBookable   int64 `json:"is_bookable,omitempty" xml:"is_bookable,omitempty"`
    

    // 是否新店：0,否；1,是
    
    IsNew   int64 `json:"is_new,omitempty" xml:"is_new,omitempty"`
    

    // 是否品牌店：0，否；1，是
    
    IsPremium   int64 `json:"is_premium,omitempty" xml:"is_premium,omitempty"`
    

    // 是否支持食安保：0，否；1，是
    
    IsInsurance   string `json:"is_insurance,omitempty" xml:"is_insurance,omitempty"`
    

    // 餐厅地址
    
    AddressText   string `json:"address_text,omitempty" xml:"address_text,omitempty"`
    

    // 餐厅营业时间
    
    ServingTimes   []string `json:"serving_times,omitempty" xml:"serving_times>string,omitempty"`
    

    // 餐厅电话
    
    PhoneList   []string `json:"phone_list,omitempty" xml:"phone_list>string,omitempty"`
    

    // 人均消费金额（元）
    
    AverageCost   string `json:"average_cost,omitempty" xml:"average_cost,omitempty"`
    

    // 餐厅整体营业状态：1 餐厅营业中，2餐厅关闭，3 餐厅网路不稳定，4 餐厅休息中，5 直接说预定，6 只接受电话预定，7 餐厅休假中
    
    TotalStatus   int64 `json:"total_status,omitempty" xml:"total_status,omitempty"`
    

    // 排序id
    
    RankId   string `json:"rank_id,omitempty" xml:"rank_id,omitempty"`
    

    // 数据列表
    
    DataList   []DataList `json:"data_list,omitempty" xml:"data_list,omitempty"`
    

    // 是否有下页
    
    HasNext   bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
    

    // 食物信息
    
    Foods   []Food `json:"foods,omitempty" xml:"foods,omitempty"`
    

    // 描述
    
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    

    // 是否为活动商品
    
    IsActivity   bool `json:"is_activity,omitempty" xml:"is_activity,omitempty"`
    

    // 分类类型. 1: 普通分类 2: 热销榜 3: 优惠分类 4: 必点分类 5: 非必点分类
    
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    

}
