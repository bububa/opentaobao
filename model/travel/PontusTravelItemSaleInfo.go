package travel

// PontusTravelItemSaleInfo 
type PontusTravelItemSaleInfo struct {

    // 减库存方式。0、拍下减库存。1、付款减库存。不传默认为0
    
    SubStock   int64 `json:"sub_stock,omitempty" xml:"sub_stock,omitempty"`
    

    // 商品售卖类型，0为日历商品，1为预约商品，2为非日历商品；默认为0
    
    SaleType   int64 `json:"sale_type,omitempty" xml:"sale_type,omitempty"`
    

    // 预约商品必填，普通商品不填。预约商品开始时间，格式：yyyy-MM-dd HH:mm
    
    BcStartDate   string `json:"bc_start_date,omitempty" xml:"bc_start_date,omitempty"`
    

    // 预约商品必填，普通商品可不填。可选出发开始日期，格式：yyyy-MM-dd。对于普通商品，根据日历库存的最早时间来自动计算。对于预约商品则为必填字段
    
    StartComboDate   string `json:"start_combo_date,omitempty" xml:"start_combo_date,omitempty"`
    

    // 预约商品必填，普通商品可不填。可选出发结束日期，格式：年-月-日 日期必须是最近300天内的，最早和最晚日期跨度最大为90天。对于普通商品，根据日历库存的最晚时间来自动计算；对于预约商品则为必填字段
    
    EndComboDate   string `json:"end_combo_date,omitempty" xml:"end_combo_date,omitempty"`
    

    // 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
    
    SecondKill   string `json:"second_kill,omitempty" xml:"second_kill,omitempty"`
    

    // 至少提前天数，最晚成团提前天数，最小0天，最大为300天；不传则为0
    
    Duration   int64 `json:"duration,omitempty" xml:"duration,omitempty"`
    

    // 最晚收客时间:小时。仅个别类目支持：境外玩乐套餐
    
    ReserveDeadlineHours   int64 `json:"reserve_deadline_hours,omitempty" xml:"reserve_deadline_hours,omitempty"`
    

    // 最晚收客时间:分钟。仅个别类目支持：境外玩乐套餐
    
    ReserveDeadlineMinutes   int64 `json:"reserve_deadline_minutes,omitempty" xml:"reserve_deadline_minutes,omitempty"`
    

    // 是否提供发票 默认为false  仅C商家需要设置该值 B商家强制提供发票
    
    HasInvoice   bool `json:"has_invoice,omitempty" xml:"has_invoice,omitempty"`
    

    // 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
    
    HasDiscount   bool `json:"has_discount,omitempty" xml:"has_discount,omitempty"`
    

    // 是否橱窗推荐，可选值：true，false；默认值：false(不推荐)
    
    HasShowcase   bool `json:"has_showcase,omitempty" xml:"has_showcase,omitempty"`
    

    // 关联商品与店铺类目 结构:"cid1,cid2,...,"。如何获取卖家店铺类目具体参见：http://open.taobao.com/doc2/apiDetail.htm?apiId=65
    
    SellerCids   []string `json:"seller_cids,omitempty" xml:"seller_cids>string,omitempty"`
    

    // 旧版电子凭证，若要支持旧版电子凭证则需填写。电子凭证码商，格式为：码商id:码商nick。 "1. 商家必须是电子凭证卖家才能发布电子凭证商品。参考链接 http://bangpai.taobao.com/group/thread/14051111-283426841.htm?spm=0.0.0.0.TRlt53 2. 发布电子凭证商品，merchant字段必填。若为淘宝发码，则merchant设置为 0:淘宝 3. network_id联系码商提供"
    
    Merchant   string `json:"merchant,omitempty" xml:"merchant,omitempty"`
    

    // 旧版电子凭证，若要支持旧版电子凭证则需填写。电子凭证网点ID
    
    NetworkId   string `json:"network_id,omitempty" xml:"network_id,omitempty"`
    

    // 旧版电子凭证，是否支持系统自动退款，true则表示支持
    
    SupportOnsaleAutoRefund   bool `json:"support_onsale_auto_refund,omitempty" xml:"support_onsale_auto_refund,omitempty"`
    

    // 新版电子凭证信息。如果传递了此参数，则sales_info中旧版电子凭证信息将被忽略，否则电子凭证信息将以旧版电子凭证参数为准。如果新、旧版参数都没传，则该商品不支持电子凭证
    
    ItemEleCertInfo   *ItemEleCertInfo `json:"item_ele_cert_info,omitempty" xml:"item_ele_cert_info,omitempty"`
    

    // 资源确认方式。1：即时确认，2：二次确认
    
    ConfirmType   int64 `json:"confirm_type,omitempty" xml:"confirm_type,omitempty"`
    

    // 资源确认时长。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认
    
    ConfirmTime   int64 `json:"confirm_time,omitempty" xml:"confirm_time,omitempty"`
    

    // 该商品真实的旅游服务提供商
    
    TouristServiceProvider   string `json:"tourist_service_provider,omitempty" xml:"tourist_service_provider,omitempty"`
    

    // 上下架状态
    
    ApproveStatus   int64 `json:"approve_status,omitempty" xml:"approve_status,omitempty"`
    

    // 返点比例
    
    AuctionPoint   int64 `json:"auction_point,omitempty" xml:"auction_point,omitempty"`
    

}
