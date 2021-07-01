package btrip

// OpenApiFlightOrderRS 
type OpenApiFlightOrderRS struct {
    // 机票订单id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 更新时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 第三方用户id
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 企业名称
    CorpName   string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
    // 商旅企业id
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 用户名称
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
    // 用户所在部门id
    DepartId   string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
    // 部门名称
    DepartName   string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
    // 商旅申请单id
    ApplyId   string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    // 联系人
    ContactName   string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
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
    // 订单状态：0待支付,1出票中,2已关闭,3有改签单,4有退票单,5出票成功,6退票申请中,7改签申请中
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 折扣
    Discount   string `json:"discount,omitempty" xml:"discount,omitempty"`
    // 航班号
    FlightNo   string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
    // 乘机人，多个用‘,’分割
    PassengerName   string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
    // 出发机场
    DepAirport   string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
    // 到达机场
    ArrAirport   string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
    // 发票信息对象
    Invoice   *OpenInvoiceDO `json:"invoice,omitempty" xml:"invoice,omitempty"`
    // 成本中心对象
    CostCenter   *OpenCostCenterDO `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
    // 价目信息
    PriceInfoList   []OpenPriceInfo `json:"price_info_list,omitempty" xml:"price_info_list>open_price_info,omitempty"`
    // 保险信息
    InsureInfoList   []OpenFlightInsureInfo `json:"insure_info_list,omitempty" xml:"insure_info_list>open_flight_insure_info,omitempty"`
    // 第三方行程id
    ThirdpartItineraryId   string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
    // 出行人列表
    UserAffiliateList   []OpenUserAffiliateDO `json:"user_affiliate_list,omitempty" xml:"user_affiliate_list>open_user_affiliate_do,omitempty"`
    // 第三方申请单ID
    ThirdpartApplyId   string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
    // 申请单名称
    BtripTitle   string `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
    // 项目id
    ProjectId   int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
    // 项目code
    ProjectCode   string `json:"project_code,omitempty" xml:"project_code,omitempty"`
    // 项目名称
    ProjectTitle   string `json:"project_title,omitempty" xml:"project_title,omitempty"`
}
