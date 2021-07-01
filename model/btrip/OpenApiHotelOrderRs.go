package btrip

// OpenApiHotelOrderRS 
type OpenApiHotelOrderRS struct {
    // 订单id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 更新时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 商旅企业id
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    // 企业名称
    CorpName   string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
    // 第三方用户id
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 用户名称
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
    // 用户所在部门id
    DepartId   string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
    // 部门名称
    DepartName   string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
    // 商旅申请单id
    ApplyId   int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    // 联系人姓名
    ContactName   string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
    // 酒店所在城市
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // 酒店名称
    HotelName   string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
    // 入住时间
    CheckIn   string `json:"check_in,omitempty" xml:"check_in,omitempty"`
    // 离店时间
    CheckOut   string `json:"check_out,omitempty" xml:"check_out,omitempty"`
    // 房型
    RoomType   string `json:"room_type,omitempty" xml:"room_type,omitempty"`
    // 房间数
    RoomNum   int64 `json:"room_num,omitempty" xml:"room_num,omitempty"`
    // 总共住几晚
    Night   int64 `json:"night,omitempty" xml:"night,omitempty"`
    // 价目详情列表
    PriceInfoList   []OpenPriceInfo `json:"price_info_list,omitempty" xml:"price_info_list>open_price_info,omitempty"`
    // 成本中心对象
    CostCenter   *OpenCostCenterDO `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
    // 入住顾客，多个用','分割
    Guest   string `json:"guest,omitempty" xml:"guest,omitempty"`
    // 订单类型描述
    OrderTypeDesc   string `json:"order_type_desc,omitempty" xml:"order_type_desc,omitempty"`
    // 订单状态描述
    OrderStatusDesc   string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
    // 订单类型
    OrderType   int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
    // 订单状态
    OrderStatus   int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
    // 发票对象
    Invoice   *OpenInvoiceDO `json:"invoice,omitempty" xml:"invoice,omitempty"`
    // 第三方行程id
    ThirdpartItineraryId   string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
    // 入住人列表
    UserAffiliateList   []OpenUserAffiliateDO `json:"user_affiliate_list,omitempty" xml:"user_affiliate_list>open_user_affiliate_do,omitempty"`
    // 第三方申请单ID
    ThirdpartApplyId   string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
    // 申请单名称
    BtripTitle   string `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
    // 项目id
    ProjectId   int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
    // 项目名称
    ProjectTitle   string `json:"project_title,omitempty" xml:"project_title,omitempty"`
    // 项目code
    ProjectCode   string `json:"project_code,omitempty" xml:"project_code,omitempty"`
}
