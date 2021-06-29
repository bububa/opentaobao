package ascpchannel

// Erpnormalconsignorderrequest 
type Erpnormalconsignorderrequest struct {
    // 发货单号（Erp单号）
    DeliveryOrderCode   string `json:"delivery_order_code,omitempty" xml:"delivery_order_code,omitempty"`
    // 店铺名称
    ShopNick   string `json:"shop_nick,omitempty" xml:"shop_nick,omitempty"`
    // 店铺id
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    // 物流订单服务类型 1:末端送装/2:干线加送装/3:仓干配
    LpServiceType   int64 `json:"lp_service_type,omitempty" xml:"lp_service_type,omitempty"`
    // 末端服务类型 1:配送 2:自提
    EndServiceType   int64 `json:"end_service_type,omitempty" xml:"end_service_type,omitempty"`
    // 订单来源平台编码,TB= 淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、PDD=拼多多、RRSLJ=日日顺乐家、MSJ=美食杰、UGO=优品购、OTHER=其他
    SourcePlatformCode   string `json:"source_platform_code,omitempty" xml:"source_platform_code,omitempty"`
    // 收件人信息
    ReceiverInfo   *Receiverinfo `json:"receiver_info,omitempty" xml:"receiver_info,omitempty"`
    // 发件人信息
    SenderInfo   *Senderinfo `json:"sender_info,omitempty" xml:"sender_info,omitempty"`
    // 退货人信息
    RefunderInfo   *Refunderinfo `json:"refunder_info,omitempty" xml:"refunder_info,omitempty"`
    // 物流公司名称
    LogisticsName   string `json:"logistics_name,omitempty" xml:"logistics_name,omitempty"`
    // 物流公司编码，SF=顺丰、EMS=标准快递、EYB=经济快件、ZJS=宅急送、YTO=圆通 、ZTO=中通 (ZTO) 、HTKY=百世汇通、UC=优速、STO=申通、TTKDEX=天天快递 、QFKD=全峰、FAST=快捷、POSTB=邮政小包 、GTO=国通、YUNDA=韵达、JD=京东配送、DD=当当宅配、OTHER=其他
    LogisticsCode   string `json:"logistics_code,omitempty" xml:"logistics_code,omitempty"`
    // 干线运单号
    ExpressCode   string `json:"express_code,omitempty" xml:"express_code,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 发货时间 格式 yyyy-MM-dd HH:mm:ss
    DeliveryTime   string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
    // 扩展字段JSON串 twoStageTms:是否二段单 0 否 ,1 是
    Feature   string `json:"feature,omitempty" xml:"feature,omitempty"`
    // 集货信息
    PickUpInfo   *Pickupinfo `json:"pick_up_info,omitempty" xml:"pick_up_info,omitempty"`
    // 发货子单信息
    ConsignOrderItemList   []Consignorderitemlist `json:"consign_order_item_list,omitempty" xml:"consign_order_item_list>consignorderitemlist,omitempty"`
    // 操作人
    Operator   string `json:"operator,omitempty" xml:"operator,omitempty"`
    // 操作时间 格式 yyyy-MM-dd HH:mm:ss
    OperateTime   string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
}