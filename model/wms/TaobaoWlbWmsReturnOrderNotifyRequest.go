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
    _orderCode   string
    // 仓库编码
    _storeCode   string
    // 用字符串格式来表示订单标记列表：比如VISIT^ SELLER_AFFORD^SYNC_RETURN_BILL 等，中间用“^”来隔开 ----------------------------------------  订单标记list（所有字母全部大写）： 9:VISIT-上门12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票
    _orderFlag   string
    // 来源单据号，销售退货时填充原发货的LBX号
    _prevOrderCode   string
    // 快递公司编码
    _tmsServiceCode   string
    // 运单号
    _tmsOrderCode   string
    // 销退时请提供退货的原因
    _returnReason   string
    // 买家昵称
    _buyerNick   string
    // 发件人信息
    _senderInfo   *Senderinfowlbwmsreturnordernotify
    // 收件人信息
    _receiverInfo   *Receiverinfowlbwmsreturnordernotify
    // 商品信息列表
    _orderItemList   []Orderitemlistwlbwmsreturnordernotify
    // 扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
    _extendFields   string
    // 备注
    _remark   string
    // 订单来源 201 淘宝，202 1688，203 苏宁，204 亚马逊中国，205 当当，206 ebay，207 唯品会，208 1号店，209 国美，210 拍拍，211 聚美优品，212 乐峰，214 京东，301 其他
    _orderSource   string
    // 订单类型 501 销退入库
    _orderType   string
    // 货主ID
    _ownerUserId   string
    // ERP订单创建时间
    _orderCreateTime   string
    // 快递公司名称
    _tmsServiceName   string
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
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderCode() string {
    return r._orderCode
}
// StoreCode Setter
// 仓库编码
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetStoreCode() string {
    return r._storeCode
}
// OrderFlag Setter
// 用字符串格式来表示订单标记列表：比如VISIT^ SELLER_AFFORD^SYNC_RETURN_BILL 等，中间用“^”来隔开 ----------------------------------------  订单标记list（所有字母全部大写）： 9:VISIT-上门12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderFlag(_orderFlag string) error {
    r._orderFlag = _orderFlag
    r.Set("order_flag", _orderFlag)
    return nil
}

// OrderFlag Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderFlag() string {
    return r._orderFlag
}
// PrevOrderCode Setter
// 来源单据号，销售退货时填充原发货的LBX号
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetPrevOrderCode(_prevOrderCode string) error {
    r._prevOrderCode = _prevOrderCode
    r.Set("prev_order_code", _prevOrderCode)
    return nil
}

// PrevOrderCode Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetPrevOrderCode() string {
    return r._prevOrderCode
}
// TmsServiceCode Setter
// 快递公司编码
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetTmsServiceCode(_tmsServiceCode string) error {
    r._tmsServiceCode = _tmsServiceCode
    r.Set("tms_service_code", _tmsServiceCode)
    return nil
}

// TmsServiceCode Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetTmsServiceCode() string {
    return r._tmsServiceCode
}
// TmsOrderCode Setter
// 运单号
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetTmsOrderCode(_tmsOrderCode string) error {
    r._tmsOrderCode = _tmsOrderCode
    r.Set("tms_order_code", _tmsOrderCode)
    return nil
}

// TmsOrderCode Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetTmsOrderCode() string {
    return r._tmsOrderCode
}
// ReturnReason Setter
// 销退时请提供退货的原因
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetReturnReason(_returnReason string) error {
    r._returnReason = _returnReason
    r.Set("return_reason", _returnReason)
    return nil
}

// ReturnReason Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetReturnReason() string {
    return r._returnReason
}
// BuyerNick Setter
// 买家昵称
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetBuyerNick() string {
    return r._buyerNick
}
// SenderInfo Setter
// 发件人信息
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetSenderInfo(_senderInfo *Senderinfowlbwmsreturnordernotify) error {
    r._senderInfo = _senderInfo
    r.Set("sender_info", _senderInfo)
    return nil
}

// SenderInfo Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetSenderInfo() *Senderinfowlbwmsreturnordernotify {
    return r._senderInfo
}
// ReceiverInfo Setter
// 收件人信息
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetReceiverInfo(_receiverInfo *Receiverinfowlbwmsreturnordernotify) error {
    r._receiverInfo = _receiverInfo
    r.Set("receiver_info", _receiverInfo)
    return nil
}

// ReceiverInfo Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetReceiverInfo() *Receiverinfowlbwmsreturnordernotify {
    return r._receiverInfo
}
// OrderItemList Setter
// 商品信息列表
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderItemList(_orderItemList []Orderitemlistwlbwmsreturnordernotify) error {
    r._orderItemList = _orderItemList
    r.Set("order_item_list", _orderItemList)
    return nil
}

// OrderItemList Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderItemList() []Orderitemlistwlbwmsreturnordernotify {
    return r._orderItemList
}
// ExtendFields Setter
// 扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetExtendFields(_extendFields string) error {
    r._extendFields = _extendFields
    r.Set("extend_fields", _extendFields)
    return nil
}

// ExtendFields Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetExtendFields() string {
    return r._extendFields
}
// Remark Setter
// 备注
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetRemark() string {
    return r._remark
}
// OrderSource Setter
// 订单来源 201 淘宝，202 1688，203 苏宁，204 亚马逊中国，205 当当，206 ebay，207 唯品会，208 1号店，209 国美，210 拍拍，211 聚美优品，212 乐峰，214 京东，301 其他
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderSource(_orderSource string) error {
    r._orderSource = _orderSource
    r.Set("order_source", _orderSource)
    return nil
}

// OrderSource Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderSource() string {
    return r._orderSource
}
// OrderType Setter
// 订单类型 501 销退入库
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderType(_orderType string) error {
    r._orderType = _orderType
    r.Set("order_type", _orderType)
    return nil
}

// OrderType Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderType() string {
    return r._orderType
}
// OwnerUserId Setter
// 货主ID
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOwnerUserId(_ownerUserId string) error {
    r._ownerUserId = _ownerUserId
    r.Set("owner_user_id", _ownerUserId)
    return nil
}

// OwnerUserId Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOwnerUserId() string {
    return r._ownerUserId
}
// OrderCreateTime Setter
// ERP订单创建时间
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderCreateTime(_orderCreateTime string) error {
    r._orderCreateTime = _orderCreateTime
    r.Set("order_create_time", _orderCreateTime)
    return nil
}

// OrderCreateTime Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderCreateTime() string {
    return r._orderCreateTime
}
// TmsServiceName Setter
// 快递公司名称
func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetTmsServiceName(_tmsServiceName string) error {
    r._tmsServiceName = _tmsServiceName
    r.Set("tms_service_name", _tmsServiceName)
    return nil
}

// TmsServiceName Getter
func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetTmsServiceName() string {
    return r._tmsServiceName
}
