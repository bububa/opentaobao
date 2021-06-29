package btrip

// OpenApiVehicleOrderRs 
type OpenApiVehicleOrderRs struct {

    // 订单id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 订单创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 订单更新时间
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    

    // 乘客名称
    
    PassengerName   string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
    

    // 商旅企业id
    
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    

    // 企业名称
    
    CorpName   string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
    

    // 预定人名称
    
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
    

    // 预定人id（第三方用户Id）
    
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

    // 用户所在部门id
    
    DeptId   int64 `json:"dept_id,omitempty" xml:"dept_id,omitempty"`
    

    // 用户所在部门名称
    
    DeptName   string `json:"dept_name,omitempty" xml:"dept_name,omitempty"`
    

    // 商旅系统审批单id
    
    ApplyId   int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    

    // 商旅审批单展示id(非审批api对接企业)
    
    ApplyShowId   string `json:"apply_show_id,omitempty" xml:"apply_show_id,omitempty"`
    

    // 真实出发地
    
    RealFromAddress   string `json:"real_from_address,omitempty" xml:"real_from_address,omitempty"`
    

    // 真实到达地
    
    RealToAddress   string `json:"real_to_address,omitempty" xml:"real_to_address,omitempty"`
    

    // 实际出发城市
    
    RealFromCityName   string `json:"real_from_city_name,omitempty" xml:"real_from_city_name,omitempty"`
    

    // 实际到达城市
    
    RealToCityName   string `json:"real_to_city_name,omitempty" xml:"real_to_city_name,omitempty"`
    

    // 出发地
    
    FromAddress   string `json:"from_address,omitempty" xml:"from_address,omitempty"`
    

    // 目的地
    
    ToAddress   string `json:"to_address,omitempty" xml:"to_address,omitempty"`
    

    // 出发城市
    
    FromCityName   string `json:"from_city_name,omitempty" xml:"from_city_name,omitempty"`
    

    // 目的城市
    
    ToCityName   string `json:"to_city_name,omitempty" xml:"to_city_name,omitempty"`
    

    // 打车事由
    
    Memo   string `json:"memo,omitempty" xml:"memo,omitempty"`
    

    // 订单状态:0:初始状态,1:已超时,2:派单成功,3:派单失败,4:已退款,5:已支付,6:已取消
    
    OrderStatus   int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
    

    // 类型级别：1经济型、2舒适型、3豪华型
    
    CarLevel   int64 `json:"car_level,omitempty" xml:"car_level,omitempty"`
    

    // 车辆类型
    
    CarInfo   string `json:"car_info,omitempty" xml:"car_info,omitempty"`
    

    // 订单预估价格
    
    EstimatePrice   string `json:"estimate_price,omitempty" xml:"estimate_price,omitempty"`
    

    // 乘客发布用车时间
    
    PublishTime   string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
    

    // 乘客上车时间
    
    TakenTime   string `json:"taken_time,omitempty" xml:"taken_time,omitempty"`
    

    // 司机到达目的地时间
    
    DriverConfirmTime   string `json:"driver_confirm_time,omitempty" xml:"driver_confirm_time,omitempty"`
    

    // 取消时间
    
    CancelTime   string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
    

    // 支付时间
    
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    

    // 行驶公里数
    
    TravelDistance   string `json:"travel_distance,omitempty" xml:"travel_distance,omitempty"`
    

    // 打车服务类型 1:出租车, 2:专车, 3:快车
    
    ServiceType   int64 `json:"service_type,omitempty" xml:"service_type,omitempty"`
    

    // 用车原因：TRAVEL: 差旅, TRAFFIC: 市内交通, WORK: 加班, OTHER: 其它
    
    BusinessCategory   string `json:"business_category,omitempty" xml:"business_category,omitempty"`
    

    // 商旅成本中心id
    
    CostCenterId   int64 `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
    

    // 成本中心编号
    
    CostCenterNumber   string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
    

    // 商旅发票id
    
    InvoiceId   int64 `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
    

    // 发票抬头
    
    InvoiceTitle   string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
    

    // 项目编号
    
    ProjectCode   string `json:"project_code,omitempty" xml:"project_code,omitempty"`
    

    // 项目名称
    
    ProjectTitle   string `json:"project_title,omitempty" xml:"project_title,omitempty"`
    

    // 价目详情列表
    
    PriceInfoList   []OpenPriceInfo `json:"price_info_list,omitempty" xml:"price_info_list,omitempty"`
    

    // 第三方行程id
    
    ThirdpartItineraryId   string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
    

    // 出行人列表
    
    UserAffiliateList   []OpenUserAffiliateDo `json:"user_affiliate_list,omitempty" xml:"user_affiliate_list,omitempty"`
    

    // 用户确认状态：0未确认；1已确认；2有异议；3系统检查不合理
    
    UserConfirm   int64 `json:"user_confirm,omitempty" xml:"user_confirm,omitempty"`
    

    // 服务商：2滴滴；3:曹操；4:首汽；5:阳光。可能会有其他服务商的增加。
    
    Provider   int64 `json:"provider,omitempty" xml:"provider,omitempty"`
    

    // 第三方申请单ID
    
    ThirdpartApplyId   string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
    

    // 申请单名称
    
    BtripTitle   string `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
    

    // 成本中心名称
    
    CostCenterName   string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
    

    // 是否特殊订单
    
    IsSpecial   bool `json:"is_special,omitempty" xml:"is_special,omitempty"`
    

    // 特殊订单类型；v_sp_t_1:用车里程，v_sp_t_2:实际下车点，v_sp_t_3:用车金额，v_sp_t_4:用车次数，v_sp_t_5:跨城订单
    
    SpecialTypes   []string `json:"special_types,omitempty" xml:"special_types>string,omitempty"`
    

    // 项目id
    
    ProjectId   int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
    

}
