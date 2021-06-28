package qimen

// DeliveryOrder 
type DeliveryOrder struct {

    // 出库单号
    
    DeliveryOrderCode   string `json:"deliveryOrderCode,omitempty" xml:"deliveryOrderCode,omitempty"`
    

    // 仓储系统出库单号
    
    DeliveryOrderId   string `json:"deliveryOrderId,omitempty" xml:"deliveryOrderId,omitempty"`
    

    // 仓库编码
    
    WarehouseCode   string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
    

    // 出库单类型(JYCK=一般交易出库;HHCK=换货出库;BFCK=补发出库;QTCK=其他出库单)
    
    OrderType   string `json:"orderType,omitempty" xml:"orderType,omitempty"`
    

    // 出库单状态(NEW-未开始处理;ACCEPT-仓库接单;PARTDELIVERED-部分发货完成;DELIVERED-发货完成;EXCEPTION-异 常;CANCELED-取消;CLOSED-关闭;REJECT-拒单;CANCELEDFAIL-取消失败;只传英文编 码)
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求 不会被重复处理;条件必填;条件为一单需要多次确认时)
    
    OutBizCode   string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
    

    // 支持出库单多次发货(多次发货后确认时;0表示发货单最终状态确认;1表示发货单中间状态确认)
    
    ConfirmType   int64 `json:"confirmType,omitempty" xml:"confirmType,omitempty"`
    

    // 订单完成时间(YYYY-MM-DD HH:MM:SS)
    
    OrderConfirmTime   string `json:"orderConfirmTime,omitempty" xml:"orderConfirmTime,omitempty"`
    

    // 当前状态操作员编码
    
    OperatorCode   string `json:"operatorCode,omitempty" xml:"operatorCode,omitempty"`
    

    // 当前状态操作员姓名
    
    OperatorName   string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
    

    // 当前状态操作时间(YYYY-MM-DD HH:MM:SS)
    
    OperateTime   string `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
    

    // 仓储费用
    
    StorageFee   string `json:"storageFee,omitempty" xml:"storageFee,omitempty"`
    

    // 发票信息
    
    Invoices   []Invoice `json:"invoices,omitempty" xml:"invoices,omitempty"`
    

    // 原出库单号(ERP分配)
    
    PreDeliveryOrderCode   string `json:"preDeliveryOrderCode,omitempty" xml:"preDeliveryOrderCode,omitempty"`
    

    // 原出库单号(WMS分配)
    
    PreDeliveryOrderId   string `json:"preDeliveryOrderId,omitempty" xml:"preDeliveryOrderId,omitempty"`
    

    // 订单标记(用字符串格式来表示订单标记列表:例如COD=货到付款;LIMIT=限时配 送;PRESELL=预 售;COMPLAIN=已投诉;SPLIT=拆单;EXCHANGE=换货;VISIT=上 门;MODIFYTRANSPORT=是否 可改配送方式;CONSIGN = 物流宝代理发货;SELLER_AFFORD=是否卖家承担运费;FENXIAO=分销订 单)
    
    OrderFlag   string `json:"orderFlag,omitempty" xml:"orderFlag,omitempty"`
    

    // 订单来源平台编码(TB=淘宝、TM=天猫、JD=京东、DD=当当、PP=拍拍、YX= 易讯、 EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、 JM=聚美、LF=乐蜂 、MGJ=蘑菇街、 JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿 里巴巴、POS=POS门店、 MIA=蜜芽、OTHER=其他(只传英文编 码))
    
    SourcePlatformCode   string `json:"sourcePlatformCode,omitempty" xml:"sourcePlatformCode,omitempty"`
    

    // 订单来源平台名称
    
    SourcePlatformName   string `json:"sourcePlatformName,omitempty" xml:"sourcePlatformName,omitempty"`
    

    // 发货单创建时间(YYYY-MM-DD HH:MM:SS)
    
    CreateTime   string `json:"createTime,omitempty" xml:"createTime,omitempty"`
    

    // 前台订单/店铺订单的创建时间/下单时间
    
    PlaceOrderTime   string `json:"placeOrderTime,omitempty" xml:"placeOrderTime,omitempty"`
    

    // 订单支付时间(YYYY-MM-DD HH:MM:SS)
    
    PayTime   string `json:"payTime,omitempty" xml:"payTime,omitempty"`
    

    // 支付平台交易号
    
    PayNo   string `json:"payNo,omitempty" xml:"payNo,omitempty"`
    

    // 店铺名称
    
    ShopNick   string `json:"shopNick,omitempty" xml:"shopNick,omitempty"`
    

    // 卖家名称
    
    SellerNick   string `json:"sellerNick,omitempty" xml:"sellerNick,omitempty"`
    

    // 买家昵称
    
    BuyerNick   string `json:"buyerNick,omitempty" xml:"buyerNick,omitempty"`
    

    // 订单总金额(订单总金额=应收金额+已收金额=商品总金额-订单折扣金额+快递费用 ;单位 元)
    
    TotalAmount   string `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
    

    // 商品总金额(元)
    
    ItemAmount   string `json:"itemAmount,omitempty" xml:"itemAmount,omitempty"`
    

    // 订单折扣金额(元)
    
    DiscountAmount   string `json:"discountAmount,omitempty" xml:"discountAmount,omitempty"`
    

    // 快递费用(元)
    
    Freight   string `json:"freight,omitempty" xml:"freight,omitempty"`
    

    // 应收金额(消费者还需要支付多少--货到付款时消费者还需要支付多少约定使用这个字 段;单位元 )
    
    ArAmount   string `json:"arAmount,omitempty" xml:"arAmount,omitempty"`
    

    // 已收金额(消费者已经支付多少;单位元)
    
    GotAmount   string `json:"gotAmount,omitempty" xml:"gotAmount,omitempty"`
    

    // COD服务费
    
    ServiceFee   string `json:"serviceFee,omitempty" xml:"serviceFee,omitempty"`
    

    // 物流公司编码(SF=顺丰、EMS=标准快递、EYB=经济快件、ZJS=宅急送、YTO=圆通 、ZTO=中 通(ZTO)、HTKY=百世汇通、UC=优速、STO=申通、TTKDEX=天天快递、QFKD=全 峰、FAST=快捷 、POSTB=邮政小包、 GTO=国通、YUNDA=韵达、JD=京东配送、DD=当当宅配、AMAZON=亚马逊物流、 OTHER=其他)
    
    LogisticsCode   string `json:"logisticsCode,omitempty" xml:"logisticsCode,omitempty"`
    

    // 物流公司名称
    
    LogisticsName   string `json:"logisticsName,omitempty" xml:"logisticsName,omitempty"`
    

    // 运单号
    
    ExpressCode   string `json:"expressCode,omitempty" xml:"expressCode,omitempty"`
    

    // 快递区域编码
    
    LogisticsAreaCode   string `json:"logisticsAreaCode,omitempty" xml:"logisticsAreaCode,omitempty"`
    

    // 发货要求
    
    DeliveryRequirements   *DeliveryRequirements `json:"deliveryRequirements,omitempty" xml:"deliveryRequirements,omitempty"`
    

    // 发件人信息
    
    SenderInfo   *SenderInfo `json:"senderInfo,omitempty" xml:"senderInfo,omitempty"`
    

    // 收件人信息
    
    ReceiverInfo   *ReceiverInfo `json:"receiverInfo,omitempty" xml:"receiverInfo,omitempty"`
    

    // 是否紧急(Y/N;默认为N)
    
    IsUrgency   string `json:"isUrgency,omitempty" xml:"isUrgency,omitempty"`
    

    // 是否需要发票(Y/N;默认为N)
    
    InvoiceFlag   string `json:"invoiceFlag,omitempty" xml:"invoiceFlag,omitempty"`
    

    // 是否需要保险(Y/N;默认为N)
    
    InsuranceFlag   string `json:"insuranceFlag,omitempty" xml:"insuranceFlag,omitempty"`
    

    // 保险信息
    
    Insurance   *Insurance `json:"insurance,omitempty" xml:"insurance,omitempty"`
    

    // 买家留言
    
    BuyerMessage   string `json:"buyerMessage,omitempty" xml:"buyerMessage,omitempty"`
    

    // 卖家留言
    
    SellerMessage   string `json:"sellerMessage,omitempty" xml:"sellerMessage,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 服务编码
    
    ServiceCode   string `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
    

    // 平台订单号
    
    OaidOrderSourceCode   string `json:"oaidOrderSourceCode,omitempty" xml:"oaidOrderSourceCode,omitempty"`
    

    // 单据列表
    
    OrderLines   []OrderLine `json:"orderLines,omitempty" xml:"orderLines,omitempty"`
    

    // 旧版本货主编码
    
    OwnerCode   string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
    

    // 最晚揽收时间, string (19) , YYYY-MM-DD HH:MM:SS
    
    LatestCollectionTime   string `json:"latestCollectionTime,omitempty" xml:"latestCollectionTime,omitempty"`
    

    // 最晚发货时间, string (19) , YYYY-MM-DD HH:MM:SS
    
    LatestDeliveryTime   string `json:"latestDeliveryTime,omitempty" xml:"latestDeliveryTime,omitempty"`
    

    // 单据总行数
    
    TotalOrderLines   int64 `json:"totalOrderLines,omitempty" xml:"totalOrderLines,omitempty"`
    

    // 该笔出库单的费用承担部门或责任部门
    
    ResponsibleDepartment   string `json:"responsibleDepartment,omitempty" xml:"responsibleDepartment,omitempty"`
    

    // 出库单确认其他出库单的子类型，用于 orderType设置为其他 出库单时设置
    
    SubOrderType   string `json:"subOrderType,omitempty" xml:"subOrderType,omitempty"`
    

    // 关联单据信息
    
    RelatedOrders   []RelatedOrder `json:"relatedOrders,omitempty" xml:"relatedOrders,omitempty"`
    

    // 要求出库时间(YYYY-MM-DD)
    
    ScheduleDate   string `json:"scheduleDate,omitempty" xml:"scheduleDate,omitempty"`
    

    // 供应商编码
    
    SupplierCode   string `json:"supplierCode,omitempty" xml:"supplierCode,omitempty"`
    

    // 供应商名称
    
    SupplierName   string `json:"supplierName,omitempty" xml:"supplierName,omitempty"`
    

    // 提货方式(到仓自提、快递、干线物流)
    
    TransportMode   string `json:"transportMode,omitempty" xml:"transportMode,omitempty"`
    

    // 提货人信息
    
    PickerInfo   *PickerInfo `json:"pickerInfo,omitempty" xml:"pickerInfo,omitempty"`
    

    // 出库单渠道类型,VIP=唯品会，FX=分销 ，SHOP=门店
    
    OrderSourceType   string `json:"orderSourceType,omitempty" xml:"orderSourceType,omitempty"`
    

    // 到货时间(YYYY-MM-DD HH:MM:SS)
    
    ReceivingTime   string `json:"receivingTime,omitempty" xml:"receivingTime,omitempty"`
    

    // 送货时间(YYYY-MM-DD HH:MM:SS)
    
    ShippingTime   string `json:"shippingTime,omitempty" xml:"shippingTime,omitempty"`
    

    // 入库仓库名称, string (50)
    
    TargetWarehouseName   string `json:"targetWarehouseName,omitempty" xml:"targetWarehouseName,omitempty"`
    

    // 入库仓库编码, string (50) ，统仓统配等无需ERP指定仓储编码的情况填OTHER
    
    TargetWarehouseCode   string `json:"targetWarehouseCode,omitempty" xml:"targetWarehouseCode,omitempty"`
    

    // 关联的入库单号（ERP分配）
    
    TargetEntryOrderCode   string `json:"targetEntryOrderCode,omitempty" xml:"targetEntryOrderCode,omitempty"`
    

    // 仓库名称
    
    WarehouseName   string `json:"warehouseName,omitempty" xml:"warehouseName,omitempty"`
    

}
