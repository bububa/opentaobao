package tmallservice

// ServiceTaskDO 
type ServiceTaskDO struct {

    // 交易实付金额
    
    ActualTotalFee   int64 `json:"actual_total_fee,omitempty" xml:"actual_total_fee,omitempty"`
    

    // 合同id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 订购关系ID
    
    OrderRelationId   int64 `json:"order_relation_id,omitempty" xml:"order_relation_id,omitempty"`
    

    // 主交易订单编号
    
    ParentBizOrderId   int64 `json:"parent_biz_order_id,omitempty" xml:"parent_biz_order_id,omitempty"`
    

    // 交易订单编号
    
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    

    // 服务订单编号
    
    ServiceOrderId   int64 `json:"service_order_id,omitempty" xml:"service_order_id,omitempty"`
    

    // 服务名称
    
    ServiceName   string `json:"service_name,omitempty" xml:"service_name,omitempty"`
    

    // 服务产品
    
    ServiceProduct   string `json:"service_product,omitempty" xml:"service_product,omitempty"`
    

    // 服务电话
    
    ServicePhone   string `json:"service_phone,omitempty" xml:"service_phone,omitempty"`
    

    // 服务次数
    
    ServiceCount   int64 `json:"service_count,omitempty" xml:"service_count,omitempty"`
    

    // 工单名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 工单描述
    
    TaskMemo   string `json:"task_memo,omitempty" xml:"task_memo,omitempty"`
    

    // 工单状态
    
    TaskStatus   int64 `json:"task_status,omitempty" xml:"task_status,omitempty"`
    

    // 工单类型
    
    TaskType   int64 `json:"task_type,omitempty" xml:"task_type,omitempty"`
    

    // 供应商编号
    
    TpId   int64 `json:"tp_id,omitempty" xml:"tp_id,omitempty"`
    

    // 服务提供商名称
    
    TpName   string `json:"tp_name,omitempty" xml:"tp_name,omitempty"`
    

    // 付款时间
    
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    

    // 付款时间
    
    PayTimeNumber   int64 `json:"pay_time_number,omitempty" xml:"pay_time_number,omitempty"`
    

    // 商品编号
    
    AuctionId   int64 `json:"auction_id,omitempty" xml:"auction_id,omitempty"`
    

    // 商品名称
    
    AuctionName   string `json:"auction_name,omitempty" xml:"auction_name,omitempty"`
    

    // 商品价格
    
    AuctionPrice   int64 `json:"auction_price,omitempty" xml:"auction_price,omitempty"`
    

    // 商品价格
    
    TaskPrice   int64 `json:"task_price,omitempty" xml:"task_price,omitempty"`
    

    // 扩展价格
    
    ExtPrice   int64 `json:"ext_price,omitempty" xml:"ext_price,omitempty"`
    

    // 型号
    
    ModelNumber   string `json:"model_number,omitempty" xml:"model_number,omitempty"`
    

    // 类目
    
    Category   string `json:"category,omitempty" xml:"category,omitempty"`
    

    // 品牌
    
    Brand   string `json:"brand,omitempty" xml:"brand,omitempty"`
    

    // 商品序列号
    
    AuctionSerialNum   string `json:"auction_serial_num,omitempty" xml:"auction_serial_num,omitempty"`
    

    // 属性
    
    Attribute   string `json:"attribute,omitempty" xml:"attribute,omitempty"`
    

    // 收货日期
    
    ReceiveTime   string `json:"receive_time,omitempty" xml:"receive_time,omitempty"`
    

    // 收货日期
    
    ReceiveTimeNumber   int64 `json:"receive_time_number,omitempty" xml:"receive_time_number,omitempty"`
    

    // 卖家id
    
    SellerId   int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
    

    // 店铺名称
    
    ShopName   string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
    

    // 卖家nick
    
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    

    // 卖家电话
    
    SellerPhone   string `json:"seller_phone,omitempty" xml:"seller_phone,omitempty"`
    

    // 卖家手机
    
    SellerMobile   string `json:"seller_mobile,omitempty" xml:"seller_mobile,omitempty"`
    

    // 买家id
    
    BuyerId   int64 `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
    

    // 买家nick
    
    BuyerNick   string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
    

    // 买家名称
    
    BuyerName   string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
    

    // 买家地址
    
    BuyerAddress   string `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
    

    // 买家邮编
    
    BuyerZipCode   string `json:"buyer_zip_code,omitempty" xml:"buyer_zip_code,omitempty"`
    

    // 买家电话
    
    BuyerPhone   string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
    

    // 买家电话
    
    BuyerMobile   string `json:"buyer_mobile,omitempty" xml:"buyer_mobile,omitempty"`
    

    // 买家邮箱
    
    BuyerMail   string `json:"buyer_mail,omitempty" xml:"buyer_mail,omitempty"`
    

    // 申请日期
    
    ApplyDate   string `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
    

    // 申请日期
    
    ApplyDateNumber   int64 `json:"apply_date_number,omitempty" xml:"apply_date_number,omitempty"`
    

    // 期望日期
    
    ExpectDate   string `json:"expect_date,omitempty" xml:"expect_date,omitempty"`
    

    // 期望日期
    
    ExpectDateNumber   int64 `json:"expect_date_number,omitempty" xml:"expect_date_number,omitempty"`
    

    // 生效日期
    
    EffectDate   string `json:"effect_date,omitempty" xml:"effect_date,omitempty"`
    

    // 生效日期
    
    EffectDateNumber   int64 `json:"effect_date_number,omitempty" xml:"effect_date_number,omitempty"`
    

    // 失效日期
    
    ExpireDate   string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
    

    // 失效日期
    
    ExpireDateNumber   int64 `json:"expire_date_number,omitempty" xml:"expire_date_number,omitempty"`
    

    // 有效期
    
    LifeCycle   int64 `json:"life_cycle,omitempty" xml:"life_cycle,omitempty"`
    

    // 异常标识
    
    Flag   int64 `json:"flag,omitempty" xml:"flag,omitempty"`
    

    // 备注信息
    
    Memo   string `json:"memo,omitempty" xml:"memo,omitempty"`
    

    // 服务商反馈信息
    
    TpFeedbackInfo   string `json:"tp_feedback_info,omitempty" xml:"tp_feedback_info,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModify   string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
    

    // 服务取消人
    
    Canceler   string `json:"canceler,omitempty" xml:"canceler,omitempty"`
    

    // 服务取消时间
    
    CancelDate   string `json:"cancel_date,omitempty" xml:"cancel_date,omitempty"`
    

    // 服务取消备注
    
    CancelMemo   string `json:"cancel_memo,omitempty" xml:"cancel_memo,omitempty"`
    

    // 合同id
    
    ContractId   int64 `json:"contract_id,omitempty" xml:"contract_id,omitempty"`
    

    // 商品以及sku的属性信息，比如计价因子长度，高度，灯头数等
    
    AuctionSkuProperties   string `json:"auction_sku_properties,omitempty" xml:"auction_sku_properties,omitempty"`
    

    // 购买的商品数量
    
    BuyAmount   int64 `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
    

    // 买家location
    
    BuyerLocation   int64 `json:"buyer_location,omitempty" xml:"buyer_location,omitempty"`
    

}
