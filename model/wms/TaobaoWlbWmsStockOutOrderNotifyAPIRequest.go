package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsStockOutOrderNotifyAPIRequest 出库单通知 API请求
// taobao.wlb.wms.stock.out.order.notify
//
// 出库单通知
type TaobaoWlbWmsStockOutOrderNotifyAPIRequest struct {
	model.Params
	// 仓储编码
	_storeCode string
	// ERP单据ID
	_orderCode string
	// 单据类型 301 调拨出库单、901普通出库单、903 其他出库单 305 B2B出库
	_orderType int64
	// ERP可选择性文本透传至WMS
	_outboundTypeDesc string
	// 订单创建时间
	_orderCreateTime string
	// 要求出库日期
	_sendTime string
	// 收件人信息
	_receiverInfo *Receiverwlbwmsstockoutordernotify
	// 发货方信息
	_senderInfo *Senderwlbwmsstockoutordernotify
	// 出库方式
	_transportMode string
	// 承运商名称
	_carriersName string
	// 取件人姓名
	_pickName string
	// 取件人电话
	_pickCall string
	// 取件人身份证ID
	_pickId string
	// 车牌号
	_carNo string
	// 订单商品信息列表
	_orderItemList []Orderitemlistwlbwmsstockoutordernotify
	// 备注
	_remark string
	// 前物流订单号，如退货入库单可能会用到
	_prevOrderCode string
	// 拓展属性
	_extendFields string
}

// NewTaobaoWlbWmsStockOutOrderNotifyRequest 初始化TaobaoWlbWmsStockOutOrderNotifyAPIRequest对象
func NewTaobaoWlbWmsStockOutOrderNotifyRequest() *TaobaoWlbWmsStockOutOrderNotifyAPIRequest {
	return &TaobaoWlbWmsStockOutOrderNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.stock.out.order.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStoreCode is StoreCode Setter
// 仓储编码
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetOrderCode is OrderCode Setter
// ERP单据ID
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetOrderType is OrderType Setter
// 单据类型 301 调拨出库单、901普通出库单、903 其他出库单 305 B2B出库
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetOrderType(_orderType int64) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetOrderType() int64 {
	return r._orderType
}

// SetOutboundTypeDesc is OutboundTypeDesc Setter
// ERP可选择性文本透传至WMS
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetOutboundTypeDesc(_outboundTypeDesc string) error {
	r._outboundTypeDesc = _outboundTypeDesc
	r.Set("outbound_type_desc", _outboundTypeDesc)
	return nil
}

// GetOutboundTypeDesc OutboundTypeDesc Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetOutboundTypeDesc() string {
	return r._outboundTypeDesc
}

// SetOrderCreateTime is OrderCreateTime Setter
// 订单创建时间
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetOrderCreateTime(_orderCreateTime string) error {
	r._orderCreateTime = _orderCreateTime
	r.Set("order_create_time", _orderCreateTime)
	return nil
}

// GetOrderCreateTime OrderCreateTime Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetOrderCreateTime() string {
	return r._orderCreateTime
}

// SetSendTime is SendTime Setter
// 要求出库日期
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetSendTime(_sendTime string) error {
	r._sendTime = _sendTime
	r.Set("send_time", _sendTime)
	return nil
}

// GetSendTime SendTime Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetSendTime() string {
	return r._sendTime
}

// SetReceiverInfo is ReceiverInfo Setter
// 收件人信息
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetReceiverInfo(_receiverInfo *Receiverwlbwmsstockoutordernotify) error {
	r._receiverInfo = _receiverInfo
	r.Set("receiver_info", _receiverInfo)
	return nil
}

// GetReceiverInfo ReceiverInfo Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetReceiverInfo() *Receiverwlbwmsstockoutordernotify {
	return r._receiverInfo
}

// SetSenderInfo is SenderInfo Setter
// 发货方信息
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetSenderInfo(_senderInfo *Senderwlbwmsstockoutordernotify) error {
	r._senderInfo = _senderInfo
	r.Set("sender_info", _senderInfo)
	return nil
}

// GetSenderInfo SenderInfo Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetSenderInfo() *Senderwlbwmsstockoutordernotify {
	return r._senderInfo
}

// SetTransportMode is TransportMode Setter
// 出库方式
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetTransportMode(_transportMode string) error {
	r._transportMode = _transportMode
	r.Set("transport_mode", _transportMode)
	return nil
}

// GetTransportMode TransportMode Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetTransportMode() string {
	return r._transportMode
}

// SetCarriersName is CarriersName Setter
// 承运商名称
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetCarriersName(_carriersName string) error {
	r._carriersName = _carriersName
	r.Set("carriers_name", _carriersName)
	return nil
}

// GetCarriersName CarriersName Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetCarriersName() string {
	return r._carriersName
}

// SetPickName is PickName Setter
// 取件人姓名
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetPickName(_pickName string) error {
	r._pickName = _pickName
	r.Set("pick_name", _pickName)
	return nil
}

// GetPickName PickName Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetPickName() string {
	return r._pickName
}

// SetPickCall is PickCall Setter
// 取件人电话
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetPickCall(_pickCall string) error {
	r._pickCall = _pickCall
	r.Set("pick_call", _pickCall)
	return nil
}

// GetPickCall PickCall Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetPickCall() string {
	return r._pickCall
}

// SetPickId is PickId Setter
// 取件人身份证ID
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetPickId(_pickId string) error {
	r._pickId = _pickId
	r.Set("pick_id", _pickId)
	return nil
}

// GetPickId PickId Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetPickId() string {
	return r._pickId
}

// SetCarNo is CarNo Setter
// 车牌号
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetCarNo(_carNo string) error {
	r._carNo = _carNo
	r.Set("car_no", _carNo)
	return nil
}

// GetCarNo CarNo Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetCarNo() string {
	return r._carNo
}

// SetOrderItemList is OrderItemList Setter
// 订单商品信息列表
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetOrderItemList(_orderItemList []Orderitemlistwlbwmsstockoutordernotify) error {
	r._orderItemList = _orderItemList
	r.Set("order_item_list", _orderItemList)
	return nil
}

// GetOrderItemList OrderItemList Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetOrderItemList() []Orderitemlistwlbwmsstockoutordernotify {
	return r._orderItemList
}

// SetRemark is Remark Setter
// 备注
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetRemark() string {
	return r._remark
}

// SetPrevOrderCode is PrevOrderCode Setter
// 前物流订单号，如退货入库单可能会用到
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetPrevOrderCode(_prevOrderCode string) error {
	r._prevOrderCode = _prevOrderCode
	r.Set("prev_order_code", _prevOrderCode)
	return nil
}

// GetPrevOrderCode PrevOrderCode Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetPrevOrderCode() string {
	return r._prevOrderCode
}

// SetExtendFields is ExtendFields Setter
// 拓展属性
func (r *TaobaoWlbWmsStockOutOrderNotifyAPIRequest) SetExtendFields(_extendFields string) error {
	r._extendFields = _extendFields
	r.Set("extend_fields", _extendFields)
	return nil
}

// GetExtendFields ExtendFields Getter
func (r TaobaoWlbWmsStockOutOrderNotifyAPIRequest) GetExtendFields() string {
	return r._extendFields
}
