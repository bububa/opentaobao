package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest 同城配送在线下单正式下单呼叫运力 API请求
// alibaba.ascp.logistics.instantsonline.calldelivery
//
// 同城配送在线下单正式下单呼叫运力
type AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest struct {
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

// NewAlibabaascplogisticsinstantsonlinecalldeliveryRequest 初始化AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest对象
func NewAlibabaascplogisticsinstantsonlinecalldeliveryRequest() *AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest {
	return &AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.instantsonline.calldelivery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemList is ItemList Setter
// 商品清单
func (r *AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) SetItemList(_itemList []ReceiverTopDto) error {
	r._itemList = _itemList
	r.Set("item_list", _itemList)
	return nil
}

// GetItemList ItemList Getter
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetItemList() []ReceiverTopDto {
	return r._itemList
}

// SetOutOrderId is OutOrderId Setter
// ERP单号
func (r *AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetBizType is BizType Setter
// 业务类型，INSTANT_ONLINE：同城配送-在线下单
func (r *AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetBizType() string {
	return r._bizType
}

// SetPickupCode is PickupCode Setter
// 指定取件码，不超过16位
func (r *AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) SetPickupCode(_pickupCode string) error {
	r._pickupCode = _pickupCode
	r.Set("pickup_code", _pickupCode)
	return nil
}

// GetPickupCode PickupCode Getter
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetPickupCode() string {
	return r._pickupCode
}

// SetTid is Tid Setter
// 淘宝订主单号
func (r *AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetTid() int64 {
	return r._tid
}

// SetSender is Sender Setter
// 发货联系人
func (r *AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) SetSender(_sender *SenderTopDto) error {
	r._sender = _sender
	r.Set("sender", _sender)
	return nil
}

// GetSender Sender Getter
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetSender() *SenderTopDto {
	return r._sender
}

// SetReceiver is Receiver Setter
// 收货联系人
func (r *AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) SetReceiver(_receiver *ReceiverTopDto) error {
	r._receiver = _receiver
	r.Set("receiver", _receiver)
	return nil
}

// GetReceiver Receiver Getter
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetReceiver() *ReceiverTopDto {
	return r._receiver
}

// SetTotalItemValue is TotalItemValue Setter
// 商品总价（原价），单位分，默认：商品单价总和
func (r *AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) SetTotalItemValue(_totalItemValue int64) error {
	r._totalItemValue = _totalItemValue
	r.Set("total_item_value", _totalItemValue)
	return nil
}

// GetTotalItemValue TotalItemValue Getter
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetTotalItemValue() int64 {
	return r._totalItemValue
}

// SetTotalItemActualValue is TotalItemActualValue Setter
// 商品实付总价（总价），单位分
func (r *AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) SetTotalItemActualValue(_totalItemActualValue int64) error {
	r._totalItemActualValue = _totalItemActualValue
	r.Set("total_item_actual_value", _totalItemActualValue)
	return nil
}

// GetTotalItemActualValue TotalItemActualValue Getter
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetTotalItemActualValue() int64 {
	return r._totalItemActualValue
}

// SetTotalWeight is TotalWeight Setter
// 总重量，单位KG，不超过20KG
func (r *AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) SetTotalWeight(_totalWeight int64) error {
	r._totalWeight = _totalWeight
	r.Set("total_weight", _totalWeight)
	return nil
}

// GetTotalWeight TotalWeight Getter
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetTotalWeight() int64 {
	return r._totalWeight
}

// SetSelectedResource is SelectedResource Setter
// 选择的物流资源
func (r *AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) SetSelectedResource(_selectedResource *SelectedResourceTopDto) error {
	r._selectedResource = _selectedResource
	r.Set("selected_resource", _selectedResource)
	return nil
}

// GetSelectedResource SelectedResource Getter
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetSelectedResource() *SelectedResourceTopDto {
	return r._selectedResource
}

// SetGenPickupCodeConfig is GenPickupCodeConfig Setter
// 生成取件码配置
func (r *AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) SetGenPickupCodeConfig(_genPickupCodeConfig *GenPickupCodeConfigTopDto) error {
	r._genPickupCodeConfig = _genPickupCodeConfig
	r.Set("gen_pickup_code_config", _genPickupCodeConfig)
	return nil
}

// GetGenPickupCodeConfig GenPickupCodeConfig Getter
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetGenPickupCodeConfig() *GenPickupCodeConfigTopDto {
	return r._genPickupCodeConfig
}

// SetDefaultTradeReceiver is DefaultTradeReceiver Setter
// 是否默认使用订单收货人地址，默认：true可不填写收货联系人
func (r *AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) SetDefaultTradeReceiver(_defaultTradeReceiver bool) error {
	r._defaultTradeReceiver = _defaultTradeReceiver
	r.Set("default_trade_receiver", _defaultTradeReceiver)
	return nil
}

// GetDefaultTradeReceiver DefaultTradeReceiver Getter
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetDefaultTradeReceiver() bool {
	return r._defaultTradeReceiver
}

// SetDefaultGenPickupCode is DefaultGenPickupCode Setter
// 默认生成取件码
func (r *AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) SetDefaultGenPickupCode(_defaultGenPickupCode bool) error {
	r._defaultGenPickupCode = _defaultGenPickupCode
	r.Set("default_gen_pickup_code", _defaultGenPickupCode)
	return nil
}

// GetDefaultGenPickupCode DefaultGenPickupCode Getter
func (r AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest) GetDefaultGenPickupCode() bool {
	return r._defaultGenPickupCode
}
