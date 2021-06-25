package wlb

// WlbOrderDetail 
type WlbOrderDetail struct {

    // 订单编码
    OrderCode   string `json:"order_code,omitempty"`

    // 对应创建物流宝订单top接口中的的out_biz_code字段，主要是用来去重用
    OrderSourceCode   string `json:"order_source_code,omitempty"`

    // 订单来源:<br/>产生物流订单的原因，比如:<br/><br/>订单来源:1:TAOBAO;2:EXT;3:ERP;4:WMS
    OrderSource   string `json:"order_source,omitempty"`

    // 仓库编码
    StoreCode   string `json:"store_code,omitempty"`

    // 1:正常订单: NARMAL<br/>2:退货订单: RETURN<br/>3:换货订单: CHANGE
    OrderType   string `json:"order_type,omitempty"`

    // (1)其它:    OTHER<br/>(2)淘宝交易: TAOBAO<br/>(3)301:调拨: ALLOCATION<br/>(4)401:盘点:CHECK<br/>(5)501:销售采购:PRUCHASE
    OrderSubType   string `json:"order_sub_type,omitempty"`

    // 出库或者入库，IN表示入库，OUT表示出库
    OperateType   string `json:"operate_type,omitempty"`

    // 物流状态，<br/>订单已创建：0<br/>订单已取消: -1<br/>订单关闭:-3<br/>下发中: 10<br/>已下发: 11<br/>接收方拒签 :-100<br/>已发货:100<br/>签收成功:200
    OrderStatus   string `json:"order_status,omitempty"`

    // 订单备注
    Remark   string `json:"remark,omitempty"`

    // 卖家ID
    UserId   int64 `json:"user_id,omitempty"`

    // 卖家NICK
    UserNick   string `json:"user_nick,omitempty"`

    // 如果是交易单，则显示交易中买家昵称
    BuyerNick   string `json:"buyer_nick,omitempty"`

    // 物流宝订单对应的商品详情
    OrderItemList   []WlbOrderItem `json:"order_item_list,omitempty"`

    // 订单创建时间
    CreateTime   string `json:"create_time,omitempty"`

    // 订单最后一次修改时间
    ModifyTime   string `json:"modify_time,omitempty"`

}
