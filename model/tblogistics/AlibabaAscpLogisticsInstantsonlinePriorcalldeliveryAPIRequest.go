package tblogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest 同城配送在线下单预询价 API请求
// alibaba.ascp.logistics.instantsonline.priorcalldelivery
//
// 同城配送在线下单预询价
type AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest struct {
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

// NewAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryRequest 初始化AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest对象
func NewAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryRequest() *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest {
	return &AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest{
		Params: model.NewParams(10),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) Reset() {
	r._itemList = r._itemList[:0]
	r._outOrderId = ""
	r._bizType = ""
	r._tid = 0
	r._sender = nil
	r._receiver = nil
	r._totalItemValue = 0
	r._totalItemActualValue = 0
	r._totalWeight = 0
	r._defaultTradeReceiver = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.instantsonline.priorcalldelivery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemList is ItemList Setter
// 商品清单
func (r *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) SetItemList(_itemList []ItemTopDto) error {
	r._itemList = _itemList
	r.Set("item_list", _itemList)
	return nil
}

// GetItemList ItemList Getter
func (r AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) GetItemList() []ItemTopDto {
	return r._itemList
}

// SetOutOrderId is OutOrderId Setter
// ERP单号
func (r *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetBizType is BizType Setter
// 业务类型，INSTANT_ONLINE：同城配送-在线下单
func (r *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) GetBizType() string {
	return r._bizType
}

// SetTid is Tid Setter
// 淘宝订主单号
func (r *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) GetTid() int64 {
	return r._tid
}

// SetSender is Sender Setter
// 发货联系人
func (r *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) SetSender(_sender *SenderTopDto) error {
	r._sender = _sender
	r.Set("sender", _sender)
	return nil
}

// GetSender Sender Getter
func (r AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) GetSender() *SenderTopDto {
	return r._sender
}

// SetReceiver is Receiver Setter
// 收货联系人
func (r *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) SetReceiver(_receiver *ReceiverTopDto) error {
	r._receiver = _receiver
	r.Set("receiver", _receiver)
	return nil
}

// GetReceiver Receiver Getter
func (r AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) GetReceiver() *ReceiverTopDto {
	return r._receiver
}

// SetTotalItemValue is TotalItemValue Setter
// 商品总价（原价），单位分，默认：商品单价总和
func (r *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) SetTotalItemValue(_totalItemValue int64) error {
	r._totalItemValue = _totalItemValue
	r.Set("total_item_value", _totalItemValue)
	return nil
}

// GetTotalItemValue TotalItemValue Getter
func (r AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) GetTotalItemValue() int64 {
	return r._totalItemValue
}

// SetTotalItemActualValue is TotalItemActualValue Setter
// 商品实付总价（总价），单位分
func (r *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) SetTotalItemActualValue(_totalItemActualValue int64) error {
	r._totalItemActualValue = _totalItemActualValue
	r.Set("total_item_actual_value", _totalItemActualValue)
	return nil
}

// GetTotalItemActualValue TotalItemActualValue Getter
func (r AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) GetTotalItemActualValue() int64 {
	return r._totalItemActualValue
}

// SetTotalWeight is TotalWeight Setter
// 总重量，单位KG，不超过20KG
func (r *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) SetTotalWeight(_totalWeight int64) error {
	r._totalWeight = _totalWeight
	r.Set("total_weight", _totalWeight)
	return nil
}

// GetTotalWeight TotalWeight Getter
func (r AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) GetTotalWeight() int64 {
	return r._totalWeight
}

// SetDefaultTradeReceiver is DefaultTradeReceiver Setter
// 是否默认使用订单收货人地址，默认：true可不填写收货联系人
func (r *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) SetDefaultTradeReceiver(_defaultTradeReceiver bool) error {
	r._defaultTradeReceiver = _defaultTradeReceiver
	r.Set("default_trade_receiver", _defaultTradeReceiver)
	return nil
}

// GetDefaultTradeReceiver DefaultTradeReceiver Getter
func (r AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) GetDefaultTradeReceiver() bool {
	return r._defaultTradeReceiver
}

var poolAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryRequest()
	},
}

// GetAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryRequest 从 sync.Pool 获取 AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest
func GetAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest() *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest {
	return poolAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest.Get().(*AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest)
}

// ReleaseAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest 将 AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest(v *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest) {
	v.Reset()
	poolAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIRequest.Put(v)
}
