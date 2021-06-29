package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
出库单通知 API请求
taobao.wlb.wms.stock.out.order.notify

出库单通知
*/
type TaobaoWlbWmsStockOutOrderNotifyRequest struct {
    model.Params
    // 仓储编码
    _storeCode   string
    // ERP单据ID
    _orderCode   string
    // 单据类型 301 调拨出库单、901普通出库单、903 其他出库单 305 B2B出库
    _orderType   int64
    // ERP可选择性文本透传至WMS
    _outboundTypeDesc   string
    // 订单创建时间
    _orderCreateTime   string
    // 要求出库日期
    _sendTime   string
    // 收件人信息
    _receiverInfo   *Receiverwlbwmsstockoutordernotify
    // 发货方信息
    _senderInfo   *Senderwlbwmsstockoutordernotify
    // 出库方式
    _transportMode   string
    // 承运商名称
    _carriersName   string
    // 取件人姓名
    _pickName   string
    // 取件人电话
    _pickCall   string
    // 取件人身份证ID
    _pickId   string
    // 车牌号
    _carNo   string
    // 订单商品信息列表
    _orderItemList   []Orderitemlistwlbwmsstockoutordernotify
    // 备注
    _remark   string
    // 前物流订单号，如退货入库单可能会用到
    _prevOrderCode   string
    // 拓展属性
    _extendFields   string
}

// 初始化TaobaoWlbWmsStockOutOrderNotifyRequest对象
func NewTaobaoWlbWmsStockOutOrderNotifyRequest() *TaobaoWlbWmsStockOutOrderNotifyRequest{
    return &TaobaoWlbWmsStockOutOrderNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.stock.out.order.notify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreCode Setter
// 仓储编码
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetStoreCode() string {
    return r._storeCode
}
// OrderCode Setter
// ERP单据ID
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetOrderCode() string {
    return r._orderCode
}
// OrderType Setter
// 单据类型 301 调拨出库单、901普通出库单、903 其他出库单 305 B2B出库
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetOrderType(_orderType int64) error {
    r._orderType = _orderType
    r.Set("order_type", _orderType)
    return nil
}

// OrderType Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetOrderType() int64 {
    return r._orderType
}
// OutboundTypeDesc Setter
// ERP可选择性文本透传至WMS
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetOutboundTypeDesc(_outboundTypeDesc string) error {
    r._outboundTypeDesc = _outboundTypeDesc
    r.Set("outbound_type_desc", _outboundTypeDesc)
    return nil
}

// OutboundTypeDesc Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetOutboundTypeDesc() string {
    return r._outboundTypeDesc
}
// OrderCreateTime Setter
// 订单创建时间
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetOrderCreateTime(_orderCreateTime string) error {
    r._orderCreateTime = _orderCreateTime
    r.Set("order_create_time", _orderCreateTime)
    return nil
}

// OrderCreateTime Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetOrderCreateTime() string {
    return r._orderCreateTime
}
// SendTime Setter
// 要求出库日期
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetSendTime(_sendTime string) error {
    r._sendTime = _sendTime
    r.Set("send_time", _sendTime)
    return nil
}

// SendTime Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetSendTime() string {
    return r._sendTime
}
// ReceiverInfo Setter
// 收件人信息
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetReceiverInfo(_receiverInfo *Receiverwlbwmsstockoutordernotify) error {
    r._receiverInfo = _receiverInfo
    r.Set("receiver_info", _receiverInfo)
    return nil
}

// ReceiverInfo Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetReceiverInfo() *Receiverwlbwmsstockoutordernotify {
    return r._receiverInfo
}
// SenderInfo Setter
// 发货方信息
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetSenderInfo(_senderInfo *Senderwlbwmsstockoutordernotify) error {
    r._senderInfo = _senderInfo
    r.Set("sender_info", _senderInfo)
    return nil
}

// SenderInfo Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetSenderInfo() *Senderwlbwmsstockoutordernotify {
    return r._senderInfo
}
// TransportMode Setter
// 出库方式
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetTransportMode(_transportMode string) error {
    r._transportMode = _transportMode
    r.Set("transport_mode", _transportMode)
    return nil
}

// TransportMode Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetTransportMode() string {
    return r._transportMode
}
// CarriersName Setter
// 承运商名称
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetCarriersName(_carriersName string) error {
    r._carriersName = _carriersName
    r.Set("carriers_name", _carriersName)
    return nil
}

// CarriersName Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetCarriersName() string {
    return r._carriersName
}
// PickName Setter
// 取件人姓名
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetPickName(_pickName string) error {
    r._pickName = _pickName
    r.Set("pick_name", _pickName)
    return nil
}

// PickName Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetPickName() string {
    return r._pickName
}
// PickCall Setter
// 取件人电话
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetPickCall(_pickCall string) error {
    r._pickCall = _pickCall
    r.Set("pick_call", _pickCall)
    return nil
}

// PickCall Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetPickCall() string {
    return r._pickCall
}
// PickId Setter
// 取件人身份证ID
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetPickId(_pickId string) error {
    r._pickId = _pickId
    r.Set("pick_id", _pickId)
    return nil
}

// PickId Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetPickId() string {
    return r._pickId
}
// CarNo Setter
// 车牌号
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetCarNo(_carNo string) error {
    r._carNo = _carNo
    r.Set("car_no", _carNo)
    return nil
}

// CarNo Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetCarNo() string {
    return r._carNo
}
// OrderItemList Setter
// 订单商品信息列表
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetOrderItemList(_orderItemList []Orderitemlistwlbwmsstockoutordernotify) error {
    r._orderItemList = _orderItemList
    r.Set("order_item_list", _orderItemList)
    return nil
}

// OrderItemList Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetOrderItemList() []Orderitemlistwlbwmsstockoutordernotify {
    return r._orderItemList
}
// Remark Setter
// 备注
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetRemark() string {
    return r._remark
}
// PrevOrderCode Setter
// 前物流订单号，如退货入库单可能会用到
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetPrevOrderCode(_prevOrderCode string) error {
    r._prevOrderCode = _prevOrderCode
    r.Set("prev_order_code", _prevOrderCode)
    return nil
}

// PrevOrderCode Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetPrevOrderCode() string {
    return r._prevOrderCode
}
// ExtendFields Setter
// 拓展属性
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetExtendFields(_extendFields string) error {
    r._extendFields = _extendFields
    r.Set("extend_fields", _extendFields)
    return nil
}

// ExtendFields Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetExtendFields() string {
    return r._extendFields
}
