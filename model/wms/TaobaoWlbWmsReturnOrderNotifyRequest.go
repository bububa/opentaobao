package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售退货通知 API请求
taobao.wlb.wms.return.order.notify

销售退货通知
*/
type TaobaoWlbWmsReturnOrderNotifyRequest struct {
    model.Params
    // ERP单据编码
    orderCode   string
    // 仓库编码
    storeCode   string
    // 用字符串格式来表示订单标记列表：比如VISIT^ SELLER_AFFORD^SYNC_RETURN_BILL 等，中间用“^”来隔开 ----------------------------------------  订单标记list（所有字母全部大写）： 9:VISIT-上门12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票
    orderFlag   string
    // 来源单据号，销售退货时填充原发货的LBX号
    prevOrderCode   string
    // 快递公司编码
    tmsServiceCode   string
    // 运单号
    tmsOrderCode   string
    // 销退时请提供退货的原因
    returnReason   string
    // 买家昵称
    buyerNick   string
    // 发件人信息
    senderInfo   *Senderinfowlbwmsreturnordernotify
    // 收件人信息
    receiverInfo   *Receiverinfowlbwmsreturnordernotify
    // 商品信息列表
    orderItemList   []Orderitemlistwlbwmsreturnordernotify
    // 扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
    extendFields   string
    // 备注
    remark   string
    // 订单来源 201 淘宝，202 1688，203 苏宁，204 亚马逊中国，205 当当，206 ebay，207 唯品会，208 1号店，209 国美，210 拍拍，211 聚美优品，212 乐峰，214 京东，301 其他
    orderSource   string
    // 订单类型 501 销退入库
    orderType   string
    // 货主ID
    ownerUserId   string
    // ERP订单创建时间
    orderCreateTime   string
    // 快递公司名称
    tmsServiceName   string
}

// 初始化TaobaoWlbWmsReturnOrderNotifyRequest对象
func NewTaobaoWlbWmsReturnOrderNotifyRequest() *TaobaoWlbWmsReturnOrderNotifyRequest{
    return &TaobaoWlbWmsReturnOrderNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.return.order.notify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// ERP单据编码
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderCode() string {
    return r.orderCode
}
// StoreCode Setter
// 仓库编码
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetStoreCode() string {
    return r.storeCode
}
// OrderFlag Setter
// 用字符串格式来表示订单标记列表：比如VISIT^ SELLER_AFFORD^SYNC_RETURN_BILL 等，中间用“^”来隔开 ----------------------------------------  订单标记list（所有字母全部大写）： 9:VISIT-上门12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderFlag(orderFlag string) error {
    r.orderFlag = orderFlag
    r.Set("order_flag", orderFlag)
    return nil
}

// OrderFlag Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderFlag() string {
    return r.orderFlag
}
// PrevOrderCode Setter
// 来源单据号，销售退货时填充原发货的LBX号
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetPrevOrderCode(prevOrderCode string) error {
    r.prevOrderCode = prevOrderCode
    r.Set("prev_order_code", prevOrderCode)
    return nil
}

// PrevOrderCode Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetPrevOrderCode() string {
    return r.prevOrderCode
}
// TmsServiceCode Setter
// 快递公司编码
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetTmsServiceCode(tmsServiceCode string) error {
    r.tmsServiceCode = tmsServiceCode
    r.Set("tms_service_code", tmsServiceCode)
    return nil
}

// TmsServiceCode Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetTmsServiceCode() string {
    return r.tmsServiceCode
}
// TmsOrderCode Setter
// 运单号
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetTmsOrderCode(tmsOrderCode string) error {
    r.tmsOrderCode = tmsOrderCode
    r.Set("tms_order_code", tmsOrderCode)
    return nil
}

// TmsOrderCode Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetTmsOrderCode() string {
    return r.tmsOrderCode
}
// ReturnReason Setter
// 销退时请提供退货的原因
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetReturnReason(returnReason string) error {
    r.returnReason = returnReason
    r.Set("return_reason", returnReason)
    return nil
}

// ReturnReason Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetReturnReason() string {
    return r.returnReason
}
// BuyerNick Setter
// 买家昵称
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetBuyerNick() string {
    return r.buyerNick
}
// SenderInfo Setter
// 发件人信息
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetSenderInfo(senderInfo *Senderinfowlbwmsreturnordernotify) error {
    r.senderInfo = senderInfo
    r.Set("sender_info", senderInfo)
    return nil
}

// SenderInfo Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetSenderInfo() *Senderinfowlbwmsreturnordernotify {
    return r.senderInfo
}
// ReceiverInfo Setter
// 收件人信息
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetReceiverInfo(receiverInfo *Receiverinfowlbwmsreturnordernotify) error {
    r.receiverInfo = receiverInfo
    r.Set("receiver_info", receiverInfo)
    return nil
}

// ReceiverInfo Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetReceiverInfo() *Receiverinfowlbwmsreturnordernotify {
    return r.receiverInfo
}
// OrderItemList Setter
// 商品信息列表
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderItemList(orderItemList []Orderitemlistwlbwmsreturnordernotify) error {
    r.orderItemList = orderItemList
    r.Set("order_item_list", orderItemList)
    return nil
}

// OrderItemList Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderItemList() []Orderitemlistwlbwmsreturnordernotify {
    return r.orderItemList
}
// ExtendFields Setter
// 扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetExtendFields(extendFields string) error {
    r.extendFields = extendFields
    r.Set("extend_fields", extendFields)
    return nil
}

// ExtendFields Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetExtendFields() string {
    return r.extendFields
}
// Remark Setter
// 备注
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

// Remark Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetRemark() string {
    return r.remark
}
// OrderSource Setter
// 订单来源 201 淘宝，202 1688，203 苏宁，204 亚马逊中国，205 当当，206 ebay，207 唯品会，208 1号店，209 国美，210 拍拍，211 聚美优品，212 乐峰，214 京东，301 其他
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderSource(orderSource string) error {
    r.orderSource = orderSource
    r.Set("order_source", orderSource)
    return nil
}

// OrderSource Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderSource() string {
    return r.orderSource
}
// OrderType Setter
// 订单类型 501 销退入库
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderType(orderType string) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

// OrderType Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderType() string {
    return r.orderType
}
// OwnerUserId Setter
// 货主ID
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOwnerUserId(ownerUserId string) error {
    r.ownerUserId = ownerUserId
    r.Set("owner_user_id", ownerUserId)
    return nil
}

// OwnerUserId Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOwnerUserId() string {
    return r.ownerUserId
}
// OrderCreateTime Setter
// ERP订单创建时间
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderCreateTime(orderCreateTime string) error {
    r.orderCreateTime = orderCreateTime
    r.Set("order_create_time", orderCreateTime)
    return nil
}

// OrderCreateTime Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderCreateTime() string {
    return r.orderCreateTime
}
// TmsServiceName Setter
// 快递公司名称
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetTmsServiceName(tmsServiceName string) error {
    r.tmsServiceName = tmsServiceName
    r.Set("tms_service_name", tmsServiceName)
    return nil
}

// TmsServiceName Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetTmsServiceName() string {
    return r.tmsServiceName
}
