package lstvending

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
售货机出货回传接口 API请求
alibaba.lst.vending.shipping.callback

零售通自动售货机商品出货回传接口，同步商品出库最新状态。
*/
type AlibabaLstVendingShippingCallbackRequest struct {
    model.Params
    // 厂商设备编码
    _equipmentCode   string
    // 交易流水号
    _tradeFlowNo   string
    // 处理结果代码
    _code   string
    // 处理结果代码描述
    _message   string
    // 出货时间
    _time   string
}

// 初始化AlibabaLstVendingShippingCallbackRequest对象
func NewAlibabaLstVendingShippingCallbackRequest() *AlibabaLstVendingShippingCallbackRequest{
    return &AlibabaLstVendingShippingCallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingShippingCallbackRequest) GetApiMethodName() string {
    return "alibaba.lst.vending.shipping.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingShippingCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EquipmentCode Setter
// 厂商设备编码
func (r *AlibabaLstVendingShippingCallbackRequest) SetEquipmentCode(_equipmentCode string) error {
    r._equipmentCode = _equipmentCode
    r.Set("equipment_code", _equipmentCode)
    return nil
}

// EquipmentCode Getter
func (r AlibabaLstVendingShippingCallbackRequest) GetEquipmentCode() string {
    return r._equipmentCode
}
// TradeFlowNo Setter
// 交易流水号
func (r *AlibabaLstVendingShippingCallbackRequest) SetTradeFlowNo(_tradeFlowNo string) error {
    r._tradeFlowNo = _tradeFlowNo
    r.Set("trade_flow_no", _tradeFlowNo)
    return nil
}

// TradeFlowNo Getter
func (r AlibabaLstVendingShippingCallbackRequest) GetTradeFlowNo() string {
    return r._tradeFlowNo
}
// Code Setter
// 处理结果代码
func (r *AlibabaLstVendingShippingCallbackRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaLstVendingShippingCallbackRequest) GetCode() string {
    return r._code
}
// Message Setter
// 处理结果代码描述
func (r *AlibabaLstVendingShippingCallbackRequest) SetMessage(_message string) error {
    r._message = _message
    r.Set("message", _message)
    return nil
}

// Message Getter
func (r AlibabaLstVendingShippingCallbackRequest) GetMessage() string {
    return r._message
}
// Time Setter
// 出货时间
func (r *AlibabaLstVendingShippingCallbackRequest) SetTime(_time string) error {
    r._time = _time
    r.Set("time", _time)
    return nil
}

// Time Getter
func (r AlibabaLstVendingShippingCallbackRequest) GetTime() string {
    return r._time
}
