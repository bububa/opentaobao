package tblogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest 同城配送在线下单正式下单呼叫运力 API请求
// alibaba.ascp.logistics.instantsonline.calldelivery
//
// 同城配送在线下单正式下单呼叫运力
type AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest struct {
	model.Params
	// 商品清单
	_itemList []ReceiverTopDto
	// ERP单号
	_outOrderId string
	// 业务类型，INSTANT_ONLINE：同城配送-在线下单
	_bizType string
	// 指定取件码，不超过16位
	_pickupCode string
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
	// 选择的物流资源
	_selectedResource *SelectedResourceTopDto
	// 生成取件码配置
	_genPickupCodeConfig *GenPickupCodeConfigTopDto
	// 是否默认使用订单收货人地址，默认：true可不填写收货联系人
	_defaultTradeReceiver bool
	// 默认生成取件码
	_defaultGenPickupCode bool
}

// NewAlibabaAscpLogisticsInstantsonlineCalldeliveryRequest 初始化AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest对象
func NewAlibabaAscpLogisticsInstantsonlineCalldeliveryRequest() *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest {
	return &AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest{
		Params: model.NewParams(14),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) Reset() {
	r._itemList = r._itemList[:0]
	r._outOrderId = ""
	r._bizType = ""
	r._pickupCode = ""
	r._tid = 0
	r._sender = nil
	r._receiver = nil
	r._totalItemValue = 0
	r._totalItemActualValue = 0
	r._totalWeight = 0
	r._selectedResource = nil
	r._genPickupCodeConfig = nil
	r._defaultTradeReceiver = false
	r._defaultGenPickupCode = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.instantsonline.calldelivery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemList is ItemList Setter
// 商品清单
func (r *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) SetItemList(_itemList []ReceiverTopDto) error {
	r._itemList = _itemList
	r.Set("item_list", _itemList)
	return nil
}

// GetItemList ItemList Getter
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetItemList() []ReceiverTopDto {
	return r._itemList
}

// SetOutOrderId is OutOrderId Setter
// ERP单号
func (r *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetBizType is BizType Setter
// 业务类型，INSTANT_ONLINE：同城配送-在线下单
func (r *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetBizType() string {
	return r._bizType
}

// SetPickupCode is PickupCode Setter
// 指定取件码，不超过16位
func (r *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) SetPickupCode(_pickupCode string) error {
	r._pickupCode = _pickupCode
	r.Set("pickup_code", _pickupCode)
	return nil
}

// GetPickupCode PickupCode Getter
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetPickupCode() string {
	return r._pickupCode
}

// SetTid is Tid Setter
// 淘宝订主单号
func (r *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetTid() int64 {
	return r._tid
}

// SetSender is Sender Setter
// 发货联系人
func (r *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) SetSender(_sender *SenderTopDto) error {
	r._sender = _sender
	r.Set("sender", _sender)
	return nil
}

// GetSender Sender Getter
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetSender() *SenderTopDto {
	return r._sender
}

// SetReceiver is Receiver Setter
// 收货联系人
func (r *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) SetReceiver(_receiver *ReceiverTopDto) error {
	r._receiver = _receiver
	r.Set("receiver", _receiver)
	return nil
}

// GetReceiver Receiver Getter
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetReceiver() *ReceiverTopDto {
	return r._receiver
}

// SetTotalItemValue is TotalItemValue Setter
// 商品总价（原价），单位分，默认：商品单价总和
func (r *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) SetTotalItemValue(_totalItemValue int64) error {
	r._totalItemValue = _totalItemValue
	r.Set("total_item_value", _totalItemValue)
	return nil
}

// GetTotalItemValue TotalItemValue Getter
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetTotalItemValue() int64 {
	return r._totalItemValue
}

// SetTotalItemActualValue is TotalItemActualValue Setter
// 商品实付总价（总价），单位分
func (r *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) SetTotalItemActualValue(_totalItemActualValue int64) error {
	r._totalItemActualValue = _totalItemActualValue
	r.Set("total_item_actual_value", _totalItemActualValue)
	return nil
}

// GetTotalItemActualValue TotalItemActualValue Getter
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetTotalItemActualValue() int64 {
	return r._totalItemActualValue
}

// SetTotalWeight is TotalWeight Setter
// 总重量，单位KG，不超过20KG
func (r *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) SetTotalWeight(_totalWeight int64) error {
	r._totalWeight = _totalWeight
	r.Set("total_weight", _totalWeight)
	return nil
}

// GetTotalWeight TotalWeight Getter
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetTotalWeight() int64 {
	return r._totalWeight
}

// SetSelectedResource is SelectedResource Setter
// 选择的物流资源
func (r *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) SetSelectedResource(_selectedResource *SelectedResourceTopDto) error {
	r._selectedResource = _selectedResource
	r.Set("selected_resource", _selectedResource)
	return nil
}

// GetSelectedResource SelectedResource Getter
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetSelectedResource() *SelectedResourceTopDto {
	return r._selectedResource
}

// SetGenPickupCodeConfig is GenPickupCodeConfig Setter
// 生成取件码配置
func (r *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) SetGenPickupCodeConfig(_genPickupCodeConfig *GenPickupCodeConfigTopDto) error {
	r._genPickupCodeConfig = _genPickupCodeConfig
	r.Set("gen_pickup_code_config", _genPickupCodeConfig)
	return nil
}

// GetGenPickupCodeConfig GenPickupCodeConfig Getter
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetGenPickupCodeConfig() *GenPickupCodeConfigTopDto {
	return r._genPickupCodeConfig
}

// SetDefaultTradeReceiver is DefaultTradeReceiver Setter
// 是否默认使用订单收货人地址，默认：true可不填写收货联系人
func (r *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) SetDefaultTradeReceiver(_defaultTradeReceiver bool) error {
	r._defaultTradeReceiver = _defaultTradeReceiver
	r.Set("default_trade_receiver", _defaultTradeReceiver)
	return nil
}

// GetDefaultTradeReceiver DefaultTradeReceiver Getter
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetDefaultTradeReceiver() bool {
	return r._defaultTradeReceiver
}

// SetDefaultGenPickupCode is DefaultGenPickupCode Setter
// 默认生成取件码
func (r *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) SetDefaultGenPickupCode(_defaultGenPickupCode bool) error {
	r._defaultGenPickupCode = _defaultGenPickupCode
	r.Set("default_gen_pickup_code", _defaultGenPickupCode)
	return nil
}

// GetDefaultGenPickupCode DefaultGenPickupCode Getter
func (r AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) GetDefaultGenPickupCode() bool {
	return r._defaultGenPickupCode
}

var poolAlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpLogisticsInstantsonlineCalldeliveryRequest()
	},
}

// GetAlibabaAscpLogisticsInstantsonlineCalldeliveryRequest 从 sync.Pool 获取 AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest
func GetAlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest() *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest {
	return poolAlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest.Get().(*AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest)
}

// ReleaseAlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest 将 AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest(v *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest) {
	v.Reset()
	poolAlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest.Put(v)
}
