package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest 同城配送在线下单预询价 API请求
// alibaba.ascp.logistics.instantsonline.priorcalldelivery
//
// 同城配送在线下单预询价
type AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest struct {
	model.Params
	// 商品清单
	_itemList []ItemTopDto
	// ERP单号
	_outOrderId string
	// 业务类型，INSTANT_ONLINE：同城配送-在线下单
	_bizType string
	// 淘宝订主单号
	_tid int64
	// 发货联系人
	_sender *SenderTopDto
	// 收货联系人
	_receiver *ReceiverTopDto
	// 商品总价（原价），单位分，默认：商品单价总和
	_totalItemValue int64
	// 商品实付总价（总价），单位分
	_totalItemActualValue int64
	// 总重量，单位KG，不超过20KG
	_totalWeight int64
	// 是否默认使用订单收货人地址，默认：true可不填写收货联系人
	_defaultTradeReceiver bool
}

// NewAlibabaascplogisticsinstantsonlinepriorcalldeliveryRequest 初始化AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest对象
func NewAlibabaascplogisticsinstantsonlinepriorcalldeliveryRequest() *AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest {
	return &AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.instantsonline.priorcalldelivery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemList is ItemList Setter
// 商品清单
func (r *AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) SetItemList(_itemList []ItemTopDto) error {
	r._itemList = _itemList
	r.Set("item_list", _itemList)
	return nil
}

// GetItemList ItemList Getter
func (r AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) GetItemList() []ItemTopDto {
	return r._itemList
}

// SetOutOrderId is OutOrderId Setter
// ERP单号
func (r *AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetBizType is BizType Setter
// 业务类型，INSTANT_ONLINE：同城配送-在线下单
func (r *AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) GetBizType() string {
	return r._bizType
}

// SetTid is Tid Setter
// 淘宝订主单号
func (r *AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) GetTid() int64 {
	return r._tid
}

// SetSender is Sender Setter
// 发货联系人
func (r *AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) SetSender(_sender *SenderTopDto) error {
	r._sender = _sender
	r.Set("sender", _sender)
	return nil
}

// GetSender Sender Getter
func (r AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) GetSender() *SenderTopDto {
	return r._sender
}

// SetReceiver is Receiver Setter
// 收货联系人
func (r *AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) SetReceiver(_receiver *ReceiverTopDto) error {
	r._receiver = _receiver
	r.Set("receiver", _receiver)
	return nil
}

// GetReceiver Receiver Getter
func (r AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) GetReceiver() *ReceiverTopDto {
	return r._receiver
}

// SetTotalItemValue is TotalItemValue Setter
// 商品总价（原价），单位分，默认：商品单价总和
func (r *AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) SetTotalItemValue(_totalItemValue int64) error {
	r._totalItemValue = _totalItemValue
	r.Set("total_item_value", _totalItemValue)
	return nil
}

// GetTotalItemValue TotalItemValue Getter
func (r AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) GetTotalItemValue() int64 {
	return r._totalItemValue
}

// SetTotalItemActualValue is TotalItemActualValue Setter
// 商品实付总价（总价），单位分
func (r *AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) SetTotalItemActualValue(_totalItemActualValue int64) error {
	r._totalItemActualValue = _totalItemActualValue
	r.Set("total_item_actual_value", _totalItemActualValue)
	return nil
}

// GetTotalItemActualValue TotalItemActualValue Getter
func (r AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) GetTotalItemActualValue() int64 {
	return r._totalItemActualValue
}

// SetTotalWeight is TotalWeight Setter
// 总重量，单位KG，不超过20KG
func (r *AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) SetTotalWeight(_totalWeight int64) error {
	r._totalWeight = _totalWeight
	r.Set("total_weight", _totalWeight)
	return nil
}

// GetTotalWeight TotalWeight Getter
func (r AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) GetTotalWeight() int64 {
	return r._totalWeight
}

// SetDefaultTradeReceiver is DefaultTradeReceiver Setter
// 是否默认使用订单收货人地址，默认：true可不填写收货联系人
func (r *AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) SetDefaultTradeReceiver(_defaultTradeReceiver bool) error {
	r._defaultTradeReceiver = _defaultTradeReceiver
	r.Set("default_trade_receiver", _defaultTradeReceiver)
	return nil
}

// GetDefaultTradeReceiver DefaultTradeReceiver Getter
func (r AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest) GetDefaultTradeReceiver() bool {
	return r._defaultTradeReceiver
}
