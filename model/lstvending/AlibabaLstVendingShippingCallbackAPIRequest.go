package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstvendingshippingcallbackAPIRequest 售货机出货回传接口 API请求
// alibaba.lst.vending.shipping.callback
//
// 零售通自动售货机商品出货回传接口，同步商品出库最新状态。
type AlibabalstvendingshippingcallbackAPIRequest struct {
	model.Params
	// 厂商设备编码
	_equipmentCode string
	// 交易流水号
	_tradeFlowNo string
	// 处理结果代码
	_code string
	// 处理结果代码描述
	_message string
	// 出货时间
	_time string
}

// NewAlibabalstvendingshippingcallbackRequest 初始化AlibabalstvendingshippingcallbackAPIRequest对象
func NewAlibabalstvendingshippingcallbackRequest() *AlibabalstvendingshippingcallbackAPIRequest {
	return &AlibabalstvendingshippingcallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstvendingshippingcallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vending.shipping.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstvendingshippingcallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstvendingshippingcallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEquipmentCode is EquipmentCode Setter
// 厂商设备编码
func (r *AlibabalstvendingshippingcallbackAPIRequest) SetEquipmentCode(_equipmentCode string) error {
	r._equipmentCode = _equipmentCode
	r.Set("equipment_code", _equipmentCode)
	return nil
}

// GetEquipmentCode EquipmentCode Getter
func (r AlibabalstvendingshippingcallbackAPIRequest) GetEquipmentCode() string {
	return r._equipmentCode
}

// SetTradeFlowNo is TradeFlowNo Setter
// 交易流水号
func (r *AlibabalstvendingshippingcallbackAPIRequest) SetTradeFlowNo(_tradeFlowNo string) error {
	r._tradeFlowNo = _tradeFlowNo
	r.Set("trade_flow_no", _tradeFlowNo)
	return nil
}

// GetTradeFlowNo TradeFlowNo Getter
func (r AlibabalstvendingshippingcallbackAPIRequest) GetTradeFlowNo() string {
	return r._tradeFlowNo
}

// SetCode is Code Setter
// 处理结果代码
func (r *AlibabalstvendingshippingcallbackAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabalstvendingshippingcallbackAPIRequest) GetCode() string {
	return r._code
}

// SetMessage is Message Setter
// 处理结果代码描述
func (r *AlibabalstvendingshippingcallbackAPIRequest) SetMessage(_message string) error {
	r._message = _message
	r.Set("message", _message)
	return nil
}

// GetMessage Message Getter
func (r AlibabalstvendingshippingcallbackAPIRequest) GetMessage() string {
	return r._message
}

// SetTime is Time Setter
// 出货时间
func (r *AlibabalstvendingshippingcallbackAPIRequest) SetTime(_time string) error {
	r._time = _time
	r.Set("time", _time)
	return nil
}

// GetTime Time Getter
func (r AlibabalstvendingshippingcallbackAPIRequest) GetTime() string {
	return r._time
}
