package alicom

// DistributionOrderInfo 
/* model for simplify = false
type DistributionOrderInfo struct {

    // 业务类型
    
    BizType   string `json:"biz_type,omitempty"`
    

    // 买家留言
    
    BuyerMessage   string `json:"buyer_message,omitempty"`
    

    // 买家名称
    
    BuyerName   string `json:"buyer_name,omitempty"`
    

    // 城市
    
    CityName   string `json:"city_name,omitempty"`
    

    // 分销商用户昵称
    
    DistributorNick   string `json:"distributor_nick,omitempty"`
    

    // 快递公司编码
    
    ExpressCode   string `json:"express_code,omitempty"`
    

    // 快递公司名称
    
    ExpressName   string `json:"express_name,omitempty"`
    

    // 快递单号
    
    ExpressNumber   string `json:"express_number,omitempty"`
    

    // 身份证件信息
    
    IdCardInfo  *struct {
        IdCardInfo  *IdCardInfo `json:"id_card_info,omitempty"`
    } `json:"id_card_info,omitempty"`
    

    // 运营商名称
    
    IspName   string `json:"isp_name,omitempty"`
    

    // 商品编号
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // 商品序列号
    
    ItemSerialNo   string `json:"item_serial_no,omitempty"`
    

    // 商品标题
    
    ItemTitle   string `json:"item_title,omitempty"`
    

    // 1-无需物流，2-供应商发货，3-分销商发货
    
    LogisticsFlag   int64 `json:"logistics_flag,omitempty"`
    

    // 物流信息,收货人信息,姓名，电话，地址
    
    LogisticsInfo  *struct {
        DistributionOrderLogisticsInfo  *DistributionOrderLogisticsInfo `json:"distribution_order_logistics_info,omitempty"`
    } `json:"logistics_info,omitempty"`
    

    // 1-未发货，2-已发货
    
    LogisticsStatus   int64 `json:"logistics_status,omitempty"`
    

    // 订购状态:NOT_ORDER-未订购，ORDER_AUDIT-订购中(无订购接口，提交给供应商，线下受理中)，ON_ORDER-订购中(有订购接口，线上受理中),SUCCESS-订购成功，FAILURE-订购失败,CANCEL-订购取消
    
    OrderStatus   string `json:"order_status,omitempty"`
    

    // 运营商合约编号
    
    OutPackageId   string `json:"out_package_id,omitempty"`
    

    // 手机号码
    
    PhoneNo   string `json:"phone_no,omitempty"`
    

    // 合约编号
    
    PlanGroupId   int64 `json:"plan_group_id,omitempty"`
    

    // 合约名称
    
    PlanGroupName   string `json:"plan_group_name,omitempty"`
    

    // 套餐编号
    
    PlanId   int64 `json:"plan_id,omitempty"`
    

    // 套餐信息
    
    PlanInfo   string `json:"plan_info,omitempty"`
    

    // 发货的商品编号，如ICCID等
    
    ProductSerialNo   string `json:"product_serial_no,omitempty"`
    

    // 省份
    
    ProvinceName   string `json:"province_name,omitempty"`
    

    // 订购失败原因
    
    Reason   string `json:"reason,omitempty"`
    

    // 备注
    
    Remarks   string `json:"remarks,omitempty"`
    

    // 订单状态
    
    Status   string `json:"status,omitempty"`
    

    // 供应商用户昵称
    
    SupplierName   string `json:"supplier_name,omitempty"`
    

    // 淘宝交易订单创建时间
    
    TbOrderCreateTime   string `json:"tb_order_create_time,omitempty"`
    

    // 淘宝交易订单号
    
    TbOrderNo   string `json:"tb_order_no,omitempty"`
    

    // 淘宝交易订单创建时间
    
    TbOrderPayTime   string `json:"tb_order_pay_time,omitempty"`
    

    // 订单价格
    
    TbOrderPrice   int64 `json:"tb_order_price,omitempty"`
    

    // 外部订单号
    
    OutOrderNo   string `json:"out_order_no,omitempty"`
    

    // 支付宝交易流水号
    
    PayOrderNo   string `json:"pay_order_no,omitempty"`
    

    // 购买数量
    
    BuyQuantity   int64 `json:"buy_quantity,omitempty"`
    

    // 套餐名称
    
    PlanName   string `json:"plan_name,omitempty"`
    

    // 号码最低消费
    
    PhoneNoMinConsume   int64 `json:"phone_no_min_consume,omitempty"`
    

    // ledgerAmount
    
    LedgerAmount   int64 `json:"ledger_amount,omitempty"`
    

    // 商品属性
    
    ItemProps   string `json:"item_props,omitempty"`
    

    // 认证的活体ID序列
    
    BiometricSeq   string `json:"biometric_seq,omitempty"`
    

}
*/

// DistributionOrderInfo 
type DistributionOrderInfo struct {

    // 业务类型
    BizType   string `json:"biz_type,omitempty"`

    // 买家留言
    BuyerMessage   string `json:"buyer_message,omitempty"`

    // 买家名称
    BuyerName   string `json:"buyer_name,omitempty"`

    // 城市
    CityName   string `json:"city_name,omitempty"`

    // 分销商用户昵称
    DistributorNick   string `json:"distributor_nick,omitempty"`

    // 快递公司编码
    ExpressCode   string `json:"express_code,omitempty"`

    // 快递公司名称
    ExpressName   string `json:"express_name,omitempty"`

    // 快递单号
    ExpressNumber   string `json:"express_number,omitempty"`

    // 身份证件信息
    IdCardInfo   *IdCardInfo `json:"id_card_info,omitempty"`

    // 运营商名称
    IspName   string `json:"isp_name,omitempty"`

    // 商品编号
    ItemId   int64 `json:"item_id,omitempty"`

    // 商品序列号
    ItemSerialNo   string `json:"item_serial_no,omitempty"`

    // 商品标题
    ItemTitle   string `json:"item_title,omitempty"`

    // 1-无需物流，2-供应商发货，3-分销商发货
    LogisticsFlag   int64 `json:"logistics_flag,omitempty"`

    // 物流信息,收货人信息,姓名，电话，地址
    LogisticsInfo   *DistributionOrderLogisticsInfo `json:"logistics_info,omitempty"`

    // 1-未发货，2-已发货
    LogisticsStatus   int64 `json:"logistics_status,omitempty"`

    // 订购状态:NOT_ORDER-未订购，ORDER_AUDIT-订购中(无订购接口，提交给供应商，线下受理中)，ON_ORDER-订购中(有订购接口，线上受理中),SUCCESS-订购成功，FAILURE-订购失败,CANCEL-订购取消
    OrderStatus   string `json:"order_status,omitempty"`

    // 运营商合约编号
    OutPackageId   string `json:"out_package_id,omitempty"`

    // 手机号码
    PhoneNo   string `json:"phone_no,omitempty"`

    // 合约编号
    PlanGroupId   int64 `json:"plan_group_id,omitempty"`

    // 合约名称
    PlanGroupName   string `json:"plan_group_name,omitempty"`

    // 套餐编号
    PlanId   int64 `json:"plan_id,omitempty"`

    // 套餐信息
    PlanInfo   string `json:"plan_info,omitempty"`

    // 发货的商品编号，如ICCID等
    ProductSerialNo   string `json:"product_serial_no,omitempty"`

    // 省份
    ProvinceName   string `json:"province_name,omitempty"`

    // 订购失败原因
    Reason   string `json:"reason,omitempty"`

    // 备注
    Remarks   string `json:"remarks,omitempty"`

    // 订单状态
    Status   string `json:"status,omitempty"`

    // 供应商用户昵称
    SupplierName   string `json:"supplier_name,omitempty"`

    // 淘宝交易订单创建时间
    TbOrderCreateTime   string `json:"tb_order_create_time,omitempty"`

    // 淘宝交易订单号
    TbOrderNo   string `json:"tb_order_no,omitempty"`

    // 淘宝交易订单创建时间
    TbOrderPayTime   string `json:"tb_order_pay_time,omitempty"`

    // 订单价格
    TbOrderPrice   int64 `json:"tb_order_price,omitempty"`

    // 外部订单号
    OutOrderNo   string `json:"out_order_no,omitempty"`

    // 支付宝交易流水号
    PayOrderNo   string `json:"pay_order_no,omitempty"`

    // 购买数量
    BuyQuantity   int64 `json:"buy_quantity,omitempty"`

    // 套餐名称
    PlanName   string `json:"plan_name,omitempty"`

    // 号码最低消费
    PhoneNoMinConsume   int64 `json:"phone_no_min_consume,omitempty"`

    // ledgerAmount
    LedgerAmount   int64 `json:"ledger_amount,omitempty"`

    // 商品属性
    ItemProps   string `json:"item_props,omitempty"`

    // 认证的活体ID序列
    BiometricSeq   string `json:"biometric_seq,omitempty"`

}
