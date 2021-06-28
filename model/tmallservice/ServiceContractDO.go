package tmallservice

// ServiceContractDO 
/* model for simplify = false
type ServiceContractDO struct {

    // 合同类型
    
    ContractType   int64 `json:"contract_type,omitempty"`
    

    // 合同id
    
    Id   int64 `json:"id,omitempty"`
    

    // 订购关系id
    
    OrderRelationId   int64 `json:"order_relation_id,omitempty"`
    

    // 主交易订单编号
    
    ParentBizOrderId   int64 `json:"parent_biz_order_id,omitempty"`
    

    // 交易订单编号
    
    BizOrderId   int64 `json:"biz_order_id,omitempty"`
    

    // 服务订单编号
    
    ServiceOrderId   int64 `json:"service_order_id,omitempty"`
    

    // 合同名称
    
    Name   string `json:"name,omitempty"`
    

    // 合同内容
    
    Content   string `json:"content,omitempty"`
    

    // 合同描述
    
    ContractMemo   string `json:"contract_memo,omitempty"`
    

    // 合同状态
    
    ContractStatus   int64 `json:"contract_status,omitempty"`
    

    // 服务名称
    
    ServiceName   string `json:"service_name,omitempty"`
    

    // 服务产品
    
    ServiceProduct   string `json:"service_product,omitempty"`
    

    // 服务电话
    
    ServicePhone   string `json:"service_phone,omitempty"`
    

    // 服务商编号
    
    TpId   int64 `json:"tp_id,omitempty"`
    

    // 供应商名称
    
    TpName   string `json:"tp_name,omitempty"`
    

    // 付款时间
    
    PayTime   string `json:"pay_time,omitempty"`
    

    // 付款时间
    
    PayTimeNumber   int64 `json:"pay_time_number,omitempty"`
    

    // 商品编号
    
    AuctionId   int64 `json:"auction_id,omitempty"`
    

    // 商品名称
    
    AuctionName   string `json:"auction_name,omitempty"`
    

    // 商品价格
    
    AuctionPrice   int64 `json:"auction_price,omitempty"`
    

    // 商品价格
    
    ContractPrice   int64 `json:"contract_price,omitempty"`
    

    // 扩展价格
    
    ExtPrice   int64 `json:"ext_price,omitempty"`
    

    // 型号
    
    ModelNumber   string `json:"model_number,omitempty"`
    

    // 类目
    
    Category   string `json:"category,omitempty"`
    

    // 品牌
    
    Brand   string `json:"brand,omitempty"`
    

    // 商品序列号
    
    AuctionSerialNum   string `json:"auction_serial_num,omitempty"`
    

    // 服务次数
    
    ServiceCount   int64 `json:"service_count,omitempty"`
    

    // 收货日期
    
    ReceiveTime   string `json:"receive_time,omitempty"`
    

    // 收货日期
    
    ReceiveTimeNumber   int64 `json:"receive_time_number,omitempty"`
    

    // 卖家id
    
    SellerId   int64 `json:"seller_id,omitempty"`
    

    // 店铺名称
    
    ShopName   string `json:"shop_name,omitempty"`
    

    // 卖家nick
    
    SellerNick   string `json:"seller_nick,omitempty"`
    

    // 卖家电话
    
    SellerPhone   string `json:"seller_phone,omitempty"`
    

    // 卖家手机
    
    SellerMobile   string `json:"seller_mobile,omitempty"`
    

    // 买家id
    
    BuyerId   int64 `json:"buyer_id,omitempty"`
    

    // 买家nick
    
    BuyerNick   string `json:"buyer_nick,omitempty"`
    

    // 买家姓名
    
    BuyerName   string `json:"buyer_name,omitempty"`
    

    // 买家地址
    
    BuyerAddress   string `json:"buyer_address,omitempty"`
    

    // 买家邮编
    
    BuyerZipCode   string `json:"buyer_zip_code,omitempty"`
    

    // 买家电话
    
    BuyerPhone   string `json:"buyer_phone,omitempty"`
    

    // 买家手机
    
    BuyerMobile   string `json:"buyer_mobile,omitempty"`
    

    // 买家有限
    
    BuyerMail   string `json:"buyer_mail,omitempty"`
    

    // 生效时间
    
    EffectDate   string `json:"effect_date,omitempty"`
    

    // 生效时间
    
    EffectDateNumber   int64 `json:"effect_date_number,omitempty"`
    

    // 失效时间
    
    ExpireDate   string `json:"expire_date,omitempty"`
    

    // 失效时间
    
    ExpireDateNumber   int64 `json:"expire_date_number,omitempty"`
    

    // 有效期
    
    LifeCycle   int64 `json:"life_cycle,omitempty"`
    

    // 电子合同编号
    
    ContractNo   int64 `json:"contract_no,omitempty"`
    

    // 电子合同url
    
    ContractUrl   string `json:"contract_url,omitempty"`
    

    // 异常标识
    
    Flag   int64 `json:"flag,omitempty"`
    

    // 备注
    
    Memo   string `json:"memo,omitempty"`
    

    // 服务反馈信息
    
    TpFeedbackInfo   string `json:"tp_feedback_info,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModify   string `json:"gmt_modify,omitempty"`
    

    // 服务取消人
    
    Canceler   string `json:"canceler,omitempty"`
    

    // 服务取消日期
    
    CancelDate   string `json:"cancel_date,omitempty"`
    

    // 服务取消备注
    
    CancelMemo   string `json:"cancel_memo,omitempty"`
    

    // 属性
    
    Attribute   string `json:"attribute,omitempty"`
    

    // 购买数量
    
    BuyAmount   int64 `json:"buy_amount,omitempty"`
    

}
*/

// ServiceContractDO 
type ServiceContractDO struct {

    // 合同类型
    ContractType   int64 `json:"contract_type,omitempty"`

    // 合同id
    Id   int64 `json:"id,omitempty"`

    // 订购关系id
    OrderRelationId   int64 `json:"order_relation_id,omitempty"`

    // 主交易订单编号
    ParentBizOrderId   int64 `json:"parent_biz_order_id,omitempty"`

    // 交易订单编号
    BizOrderId   int64 `json:"biz_order_id,omitempty"`

    // 服务订单编号
    ServiceOrderId   int64 `json:"service_order_id,omitempty"`

    // 合同名称
    Name   string `json:"name,omitempty"`

    // 合同内容
    Content   string `json:"content,omitempty"`

    // 合同描述
    ContractMemo   string `json:"contract_memo,omitempty"`

    // 合同状态
    ContractStatus   int64 `json:"contract_status,omitempty"`

    // 服务名称
    ServiceName   string `json:"service_name,omitempty"`

    // 服务产品
    ServiceProduct   string `json:"service_product,omitempty"`

    // 服务电话
    ServicePhone   string `json:"service_phone,omitempty"`

    // 服务商编号
    TpId   int64 `json:"tp_id,omitempty"`

    // 供应商名称
    TpName   string `json:"tp_name,omitempty"`

    // 付款时间
    PayTime   string `json:"pay_time,omitempty"`

    // 付款时间
    PayTimeNumber   int64 `json:"pay_time_number,omitempty"`

    // 商品编号
    AuctionId   int64 `json:"auction_id,omitempty"`

    // 商品名称
    AuctionName   string `json:"auction_name,omitempty"`

    // 商品价格
    AuctionPrice   int64 `json:"auction_price,omitempty"`

    // 商品价格
    ContractPrice   int64 `json:"contract_price,omitempty"`

    // 扩展价格
    ExtPrice   int64 `json:"ext_price,omitempty"`

    // 型号
    ModelNumber   string `json:"model_number,omitempty"`

    // 类目
    Category   string `json:"category,omitempty"`

    // 品牌
    Brand   string `json:"brand,omitempty"`

    // 商品序列号
    AuctionSerialNum   string `json:"auction_serial_num,omitempty"`

    // 服务次数
    ServiceCount   int64 `json:"service_count,omitempty"`

    // 收货日期
    ReceiveTime   string `json:"receive_time,omitempty"`

    // 收货日期
    ReceiveTimeNumber   int64 `json:"receive_time_number,omitempty"`

    // 卖家id
    SellerId   int64 `json:"seller_id,omitempty"`

    // 店铺名称
    ShopName   string `json:"shop_name,omitempty"`

    // 卖家nick
    SellerNick   string `json:"seller_nick,omitempty"`

    // 卖家电话
    SellerPhone   string `json:"seller_phone,omitempty"`

    // 卖家手机
    SellerMobile   string `json:"seller_mobile,omitempty"`

    // 买家id
    BuyerId   int64 `json:"buyer_id,omitempty"`

    // 买家nick
    BuyerNick   string `json:"buyer_nick,omitempty"`

    // 买家姓名
    BuyerName   string `json:"buyer_name,omitempty"`

    // 买家地址
    BuyerAddress   string `json:"buyer_address,omitempty"`

    // 买家邮编
    BuyerZipCode   string `json:"buyer_zip_code,omitempty"`

    // 买家电话
    BuyerPhone   string `json:"buyer_phone,omitempty"`

    // 买家手机
    BuyerMobile   string `json:"buyer_mobile,omitempty"`

    // 买家有限
    BuyerMail   string `json:"buyer_mail,omitempty"`

    // 生效时间
    EffectDate   string `json:"effect_date,omitempty"`

    // 生效时间
    EffectDateNumber   int64 `json:"effect_date_number,omitempty"`

    // 失效时间
    ExpireDate   string `json:"expire_date,omitempty"`

    // 失效时间
    ExpireDateNumber   int64 `json:"expire_date_number,omitempty"`

    // 有效期
    LifeCycle   int64 `json:"life_cycle,omitempty"`

    // 电子合同编号
    ContractNo   int64 `json:"contract_no,omitempty"`

    // 电子合同url
    ContractUrl   string `json:"contract_url,omitempty"`

    // 异常标识
    Flag   int64 `json:"flag,omitempty"`

    // 备注
    Memo   string `json:"memo,omitempty"`

    // 服务反馈信息
    TpFeedbackInfo   string `json:"tp_feedback_info,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModify   string `json:"gmt_modify,omitempty"`

    // 服务取消人
    Canceler   string `json:"canceler,omitempty"`

    // 服务取消日期
    CancelDate   string `json:"cancel_date,omitempty"`

    // 服务取消备注
    CancelMemo   string `json:"cancel_memo,omitempty"`

    // 属性
    Attribute   string `json:"attribute,omitempty"`

    // 购买数量
    BuyAmount   int64 `json:"buy_amount,omitempty"`

}
