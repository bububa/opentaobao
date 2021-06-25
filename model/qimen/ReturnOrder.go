package qimen

// ReturnOrder 
type ReturnOrder struct {

    // ERP的退货入库单编码
    ReturnOrderCode   string `json:"returnOrderCode,omitempty"`

    // 仓库系统订单编码
    ReturnOrderId   string `json:"returnOrderId,omitempty"`

    // 仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // 外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求不会 被重复处理)
    OutBizCode   string `json:"outBizCode,omitempty"`

    // 单据类型(THRK=退货入库;HHRK=换货入库;只传英文编码)
    OrderType   string `json:"orderType,omitempty"`

    // 确认入库时间(YYYY-MM-DD HH:MM:SS)
    OrderConfirmTime   string `json:"orderConfirmTime,omitempty"`

    // 物流公司编码(SF=顺丰、EMS=标准快递、EYB=经济快件、ZJS=宅急送、YTO=圆通、ZTO=中通(ZTO)、HTKY=百世汇通、 UC=优速、STO=申通、TTKDEX=天天快递、QFKD=全峰、FAST=快捷、POSTB=邮政小包、GTO=国通、YUNDA=韵达、JD=京东配送、DD=当当宅配、 AMAZON=亚马逊物流、OTHER=其他;只传英文编码)
    LogisticsCode   string `json:"logisticsCode,omitempty"`

    // 物流公司名称
    LogisticsName   string `json:"logisticsName,omitempty"`

    // 运单号
    ExpressCode   string `json:"expressCode,omitempty"`

    // 退货原因
    ReturnReason   string `json:"returnReason,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 发件人信息
    SenderInfo   *SenderInfo `json:"senderInfo,omitempty"`

    // 用字符串格式来表示订单标记列表(比如VISIT^ SELLER_AFFORD^SYNC_RETURN_BILL等;中间用“^”来隔开订单标记list (所 有字母全部大写) VISIT=上门；SELLER_AFFORD=是否卖家承担运费(默认是)SYNC_RETURN_BILL=同时退回发票)
    OrderFlag   string `json:"orderFlag,omitempty"`

    // 原出库单号(ERP分配)
    PreDeliveryOrderCode   string `json:"preDeliveryOrderCode,omitempty"`

    // 原出库单号(WMS分配)
    PreDeliveryOrderId   string `json:"preDeliveryOrderId,omitempty"`

    // 买家昵称
    BuyerNick   string `json:"buyerNick,omitempty"`

    // 订单来源平台编码, string (50),TB= 淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、MIA=蜜芽、GW=商家官网、CT=村淘、YJWD=云集微店、PDD=拼多多、OTHERS=其他,
    SourcePlatformCode   string `json:"sourcePlatformCode,omitempty"`

    // 订单来源平台名称
    SourcePlatformName   string `json:"sourcePlatformName,omitempty"`

    // 店铺名称
    ShopNick   string `json:"shopNick,omitempty"`

    // 卖家名称
    SellerNick   string `json:"sellerNick,omitempty"`

}
