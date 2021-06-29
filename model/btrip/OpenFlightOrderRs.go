package btrip

// OpenFlightOrderRs 
type OpenFlightOrderRs struct {

    // 机票订单id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 更新时间
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    

    // 企业id
    
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    

    // 企业名称
    
    CorpName   string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
    

    // 用户id
    
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

    // 用户名称
    
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
    

    // 部门id
    
    DepartId   string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
    

    // 部门名称
    
    DepartName   string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
    

    // 申请单id
    
    ApplyId   int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    

    // 联系人名称
    
    ContactName   string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
    

    // 联系人电话
    
    ContactPhone   string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
    

    // 出发城市
    
    DepCity   string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
    

    // 到达城市
    
    ArrCity   string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
    

    // 出发日期
    
    DepDate   string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
    

    // 到达日期
    
    RetDate   string `json:"ret_date,omitempty" xml:"ret_date,omitempty"`
    

    // 行程类型。0:单程，1:往返，2:中转
    
    TripType   int64 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
    

    // 乘机人数量
    
    PassengerCount   int64 `json:"passenger_count,omitempty" xml:"passenger_count,omitempty"`
    

    // 舱位类型
    
    CabinClass   string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
    

    // 订单状态：0待支付,1出票中,2已关闭(如：未支付已关闭),3有改签单,4有退票单,5出票成功,6退票申请中,7改签申请中
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 保险信息
    
    InsureInfoList   []OpenFlightInsureInfo `json:"insure_info_list,omitempty" xml:"insure_info_list,omitempty"`
    

    // 价目信息
    
    PriceInfoList   []OpenPriceInfo `json:"price_info_list,omitempty" xml:"price_info_list,omitempty"`
    

    // 成本中心对象
    
    CostCenter   *OpenCostCenterDo `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
    

    // invoiceDO
    
    Invoice   *OpenInvoiceDo `json:"invoice,omitempty" xml:"invoice,omitempty"`
    

    // 到达机场
    
    ArrAirport   string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
    

    // 出发机场
    
    DepAirport   string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
    

    // 乘机人，多个用‘,’分割
    
    PassengerName   string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
    

    // 航班号
    
    FlightNo   string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
    

    // 折扣
    
    Discount   string `json:"discount,omitempty" xml:"discount,omitempty"`
    

    // 第三方行程id
    
    ThirdpartItineraryId   string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
    

    // 出行人列表
    
    UserAffiliateList   []OpenUserAffiliateDo `json:"user_affiliate_list,omitempty" xml:"user_affiliate_list,omitempty"`
    

}
