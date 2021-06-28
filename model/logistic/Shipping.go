package logistic

// Shipping 
/* model for simplify = false
type Shipping struct {

    // 返回发货是否成功。
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 交易ID
    
    Tid   int64 `json:"tid,omitempty"`
    

    // 物流订单编号
    
    OrderCode   string `json:"order_code,omitempty"`
    

    // 物流订单状态,可选值:CREATED(订单已创建) RECREATED(订单重新创建) CANCELLED(订单已取消) CLOSED(订单关闭) SENDING(等候发送给物流公司) ACCEPTING(已发送给物流公司,等待接单) ACCEPTED(物流公司已接单) REJECTED(物流公司不接单) PICK_UP(物流公司揽收成功) PICK_UP_FAILED(物流公司揽收失败) LOST(物流公司丢单) REJECTED_BY_RECEIVER(对方拒签) ACCEPTED_BY_RECEIVER(发货方式在线下单：对方已签收；自己联系：卖家已发货)
    
    Status   string `json:"status,omitempty"`
    

    // 标示为是否快捷COD订单
    
    IsQuickCodOrder   bool `json:"is_quick_cod_order,omitempty"`
    

    // 卖家昵称
    
    SellerNick   string `json:"seller_nick,omitempty"`
    

    // 买家昵称
    
    BuyerNick   string `json:"buyer_nick,omitempty"`
    

    // 预约取货开始时间
    
    DeliveryStart   string `json:"delivery_start,omitempty"`
    

    // 预约取货结束时间
    
    DeliveryEnd   string `json:"delivery_end,omitempty"`
    

    // 运单号.具体一个物流公司的运单号码.
    
    OutSid   string `json:"out_sid,omitempty"`
    

    // 货物名称
    
    ItemTitle   string `json:"item_title,omitempty"`
    

    // 收件人姓名
    
    ReceiverName   string `json:"receiver_name,omitempty"`
    

    // 收件人电话
    
    ReceiverPhone   string `json:"receiver_phone,omitempty"`
    

    // 收件人手机号码
    
    ReceiverMobile   string `json:"receiver_mobile,omitempty"`
    

    // 收件人地址信息(在传输请求参数Fields字段时，必须使用“receiver_location”才能返回此字段)
    
    Location  *struct {
        Location  *Location `json:"location,omitempty"`
    } `json:"location,omitempty"`
    

    // 物流方式.可选值:free(卖家包邮),post(平邮),express(快递),ems(EMS).
    
    Type   string `json:"type,omitempty"`
    

    // 谁承担运费.可选值:buyer(买家承担),seller(卖家承担运费).
    
    FreightPayer   string `json:"freight_payer,omitempty"`
    

    // 卖家是否确认发货.可选值:yes(是),no(否).
    
    SellerConfirm   string `json:"seller_confirm,omitempty"`
    

    // 物流公司名称
    
    CompanyName   string `json:"company_name,omitempty"`
    

    // 运单创建时间
    
    Created   string `json:"created,omitempty"`
    

    // 运单修改时间
    
    Modified   string `json:"modified,omitempty"`
    

    // 拆单子订单列表，对应的数据是：该物流订单下的全部子订单
    
    SubTids  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"sub_tids,omitempty"`
    

    // 表明是否是拆单，默认值0，1表示拆单
    
    IsSplit   int64 `json:"is_split,omitempty"`
    

}
*/

// Shipping 
type Shipping struct {

    // 返回发货是否成功。
    IsSuccess   bool `json:"is_success,omitempty"`

    // 交易ID
    Tid   int64 `json:"tid,omitempty"`

    // 物流订单编号
    OrderCode   string `json:"order_code,omitempty"`

    // 物流订单状态,可选值:CREATED(订单已创建) RECREATED(订单重新创建) CANCELLED(订单已取消) CLOSED(订单关闭) SENDING(等候发送给物流公司) ACCEPTING(已发送给物流公司,等待接单) ACCEPTED(物流公司已接单) REJECTED(物流公司不接单) PICK_UP(物流公司揽收成功) PICK_UP_FAILED(物流公司揽收失败) LOST(物流公司丢单) REJECTED_BY_RECEIVER(对方拒签) ACCEPTED_BY_RECEIVER(发货方式在线下单：对方已签收；自己联系：卖家已发货)
    Status   string `json:"status,omitempty"`

    // 标示为是否快捷COD订单
    IsQuickCodOrder   bool `json:"is_quick_cod_order,omitempty"`

    // 卖家昵称
    SellerNick   string `json:"seller_nick,omitempty"`

    // 买家昵称
    BuyerNick   string `json:"buyer_nick,omitempty"`

    // 预约取货开始时间
    DeliveryStart   string `json:"delivery_start,omitempty"`

    // 预约取货结束时间
    DeliveryEnd   string `json:"delivery_end,omitempty"`

    // 运单号.具体一个物流公司的运单号码.
    OutSid   string `json:"out_sid,omitempty"`

    // 货物名称
    ItemTitle   string `json:"item_title,omitempty"`

    // 收件人姓名
    ReceiverName   string `json:"receiver_name,omitempty"`

    // 收件人电话
    ReceiverPhone   string `json:"receiver_phone,omitempty"`

    // 收件人手机号码
    ReceiverMobile   string `json:"receiver_mobile,omitempty"`

    // 收件人地址信息(在传输请求参数Fields字段时，必须使用“receiver_location”才能返回此字段)
    Location   *Location `json:"location,omitempty"`

    // 物流方式.可选值:free(卖家包邮),post(平邮),express(快递),ems(EMS).
    Type   string `json:"type,omitempty"`

    // 谁承担运费.可选值:buyer(买家承担),seller(卖家承担运费).
    FreightPayer   string `json:"freight_payer,omitempty"`

    // 卖家是否确认发货.可选值:yes(是),no(否).
    SellerConfirm   string `json:"seller_confirm,omitempty"`

    // 物流公司名称
    CompanyName   string `json:"company_name,omitempty"`

    // 运单创建时间
    Created   string `json:"created,omitempty"`

    // 运单修改时间
    Modified   string `json:"modified,omitempty"`

    // 拆单子订单列表，对应的数据是：该物流订单下的全部子订单
    SubTids   []int64 `json:"sub_tids,omitempty"`

    // 表明是否是拆单，默认值0，1表示拆单
    IsSplit   int64 `json:"is_split,omitempty"`

}
