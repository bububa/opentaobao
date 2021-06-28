package cntms

// CnTmsLogisticsOrderConsignContent 
/* model for simplify = false
type CnTmsLogisticsOrderConsignContent struct {

    // ERP订单号
    
    OrderCode   string `json:"order_code,omitempty"`
    

    // 交易订单id或者商家订单号； 若阿里系订单，必须与阿里对应
    
    TradeId   string `json:"trade_id,omitempty"`
    

    // 来源渠道（TB 淘宝，JD 京东，TM 天猫，1688 1688（阿里中文站），YHD 1号店，DD 当当，VANCL 凡客，PP 拍拍，YX 易讯，EBAY 易贝ebay，AMAZON 亚马逊，SN 苏宁在线，GM 国美在线，WPH 唯品会，JM 聚美优品，LF 乐蜂网，MGJ 蘑菇街，JS 聚尚网，YG 优购，YT 银泰，YL 邮乐，PX 拍鞋网，POS POS门店，OTHERS 其他）
    
    OrderSource   string `json:"order_source,omitempty"`
    

    // 店铺编码
    
    ShopCode   string `json:"shop_code,omitempty"`
    

    // 运单号
    
    MailNo   string `json:"mail_no,omitempty"`
    

    // 物流公司编码
    
    TmsCode   string `json:"tms_code,omitempty"`
    

    // 配送发货单收件人信息
    
    ReceiverInfo  *struct {
        CnTmsLogisticsOrderReceiverInfo  *CnTmsLogisticsOrderReceiverInfo `json:"cn_tms_logistics_order_receiver_info,omitempty"`
    } `json:"receiver_info,omitempty"`
    

    // 配送发货单发件人信息
    
    SenderInfo  *struct {
        CnTmsLogisticsOrderSenderinfo  *CnTmsLogisticsOrderSenderinfo `json:"cn_tms_logistics_order_senderinfo,omitempty"`
    } `json:"sender_info,omitempty"`
    

    // 发货商品信息，最大50条记录
    
    Items  struct {
        CnTmsLogisticsOrderItem  []CnTmsLogisticsOrderItem `json:"cn_tms_logistics_order_item,omitempty"`
    } `json:"items,omitempty"`
    

    // 配送要求信息（当前业务暂不支持）
    
    DeliverRequirements  *struct {
        CnTmsLogisticsOrderDeliverRequirements  *CnTmsLogisticsOrderDeliverRequirements `json:"cn_tms_logistics_order_deliver_requirements,omitempty"`
    } `json:"deliver_requirements,omitempty"`
    

    // 商家送货方式，1商家送货，2菜鸟揽货
    
    PickUpType   int64 `json:"pick_up_type,omitempty"`
    

    // 要求菜鸟上门揽货服务，当pick_up_Type=2且需求指定时做揽收时，此字段需要传值（当前业务暂不支持）
    
    TmsGotService  *struct {
        CnTmsLogisticsOrderGotService  *CnTmsLogisticsOrderGotService `json:"cn_tms_logistics_order_got_service,omitempty"`
    } `json:"tms_got_service,omitempty"`
    

    // 物流服务解决方案Code，此字段由菜鸟提供
    
    SolutionsCode   string `json:"solutions_code,omitempty"`
    

    // 此订单的第几个包裹，如订单拆包裹时，传入此参数，配送时会将同一订单的包裹一配送
    
    PackageNo   int64 `json:"package_no,omitempty"`
    

    // 包裹重量（克）
    
    PackageWeight   int64 `json:"package_weight,omitempty"`
    

    // 此订单总的包裹数，如订单拆包裹时，传入此参数，配送时会将同一订单的包裹一配送
    
    PackageCount   int64 `json:"package_count,omitempty"`
    

    // 包裹长度（厘米）
    
    PackageLength   int64 `json:"package_length,omitempty"`
    

    // 扩展字段 K:V;
    
    ExtendFields   string `json:"extend_fields,omitempty"`
    

    // 包裹高度（厘米）
    
    PackageHeight   int64 `json:"package_height,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

    // 包裹体积（立方厘米）
    
    PackageVolume   int64 `json:"package_volume,omitempty"`
    

    // 包裹宽度（厘米）
    
    PackageWidth   int64 `json:"package_width,omitempty"`
    

    // 关单标示，true表示发货完结
    
    IsLastBatch   bool `json:"is_last_batch,omitempty"`
    

    // 包裹列表，支持一单多包裹
    
    PackageList  struct {
        CnTmsLogisticsOrderItemPackageInfo  []CnTmsLogisticsOrderItemPackageInfo `json:"cn_tms_logistics_order_item_package_info,omitempty"`
    } `json:"package_list,omitempty"`
    

}
*/

// CnTmsLogisticsOrderConsignContent 
type CnTmsLogisticsOrderConsignContent struct {

    // ERP订单号
    OrderCode   string `json:"order_code,omitempty"`

    // 交易订单id或者商家订单号； 若阿里系订单，必须与阿里对应
    TradeId   string `json:"trade_id,omitempty"`

    // 来源渠道（TB 淘宝，JD 京东，TM 天猫，1688 1688（阿里中文站），YHD 1号店，DD 当当，VANCL 凡客，PP 拍拍，YX 易讯，EBAY 易贝ebay，AMAZON 亚马逊，SN 苏宁在线，GM 国美在线，WPH 唯品会，JM 聚美优品，LF 乐蜂网，MGJ 蘑菇街，JS 聚尚网，YG 优购，YT 银泰，YL 邮乐，PX 拍鞋网，POS POS门店，OTHERS 其他）
    OrderSource   string `json:"order_source,omitempty"`

    // 店铺编码
    ShopCode   string `json:"shop_code,omitempty"`

    // 运单号
    MailNo   string `json:"mail_no,omitempty"`

    // 物流公司编码
    TmsCode   string `json:"tms_code,omitempty"`

    // 配送发货单收件人信息
    ReceiverInfo   *CnTmsLogisticsOrderReceiverInfo `json:"receiver_info,omitempty"`

    // 配送发货单发件人信息
    SenderInfo   *CnTmsLogisticsOrderSenderinfo `json:"sender_info,omitempty"`

    // 发货商品信息，最大50条记录
    Items   []CnTmsLogisticsOrderItem `json:"items,omitempty"`

    // 配送要求信息（当前业务暂不支持）
    DeliverRequirements   *CnTmsLogisticsOrderDeliverRequirements `json:"deliver_requirements,omitempty"`

    // 商家送货方式，1商家送货，2菜鸟揽货
    PickUpType   int64 `json:"pick_up_type,omitempty"`

    // 要求菜鸟上门揽货服务，当pick_up_Type=2且需求指定时做揽收时，此字段需要传值（当前业务暂不支持）
    TmsGotService   *CnTmsLogisticsOrderGotService `json:"tms_got_service,omitempty"`

    // 物流服务解决方案Code，此字段由菜鸟提供
    SolutionsCode   string `json:"solutions_code,omitempty"`

    // 此订单的第几个包裹，如订单拆包裹时，传入此参数，配送时会将同一订单的包裹一配送
    PackageNo   int64 `json:"package_no,omitempty"`

    // 包裹重量（克）
    PackageWeight   int64 `json:"package_weight,omitempty"`

    // 此订单总的包裹数，如订单拆包裹时，传入此参数，配送时会将同一订单的包裹一配送
    PackageCount   int64 `json:"package_count,omitempty"`

    // 包裹长度（厘米）
    PackageLength   int64 `json:"package_length,omitempty"`

    // 扩展字段 K:V;
    ExtendFields   string `json:"extend_fields,omitempty"`

    // 包裹高度（厘米）
    PackageHeight   int64 `json:"package_height,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 包裹体积（立方厘米）
    PackageVolume   int64 `json:"package_volume,omitempty"`

    // 包裹宽度（厘米）
    PackageWidth   int64 `json:"package_width,omitempty"`

    // 关单标示，true表示发货完结
    IsLastBatch   bool `json:"is_last_batch,omitempty"`

    // 包裹列表，支持一单多包裹
    PackageList   []CnTmsLogisticsOrderItemPackageInfo `json:"package_list,omitempty"`

}
