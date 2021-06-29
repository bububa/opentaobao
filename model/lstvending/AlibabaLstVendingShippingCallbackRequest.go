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
    equipmentCode   string
    // 交易流水号
    tradeFlowNo   string
    // 处理结果代码
    code   string
    // 处理结果代码描述
    message   string
    // 出货时间
    time   string
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
func (r *AlibabaLstVendingShippingCallbackRequest) SetEquipmentCode(equipmentCode string) error {
    r.equipmentCode = equipmentCode
    r.Set("equipment_code", equipmentCode)
    return nil
}

// EquipmentCode Getter
func (r AlibabaLstVendingShippingCallbackRequest) GetEquipmentCode() string {
    return r.equipmentCode
}
// TradeFlowNo Setter
// 交易流水号
func (r *AlibabaLstVendingShippingCallbackRequest) SetTradeFlowNo(tradeFlowNo string) error {
    r.tradeFlowNo = tradeFlowNo
    r.Set("trade_flow_no", tradeFlowNo)
    return nil
}

// TradeFlowNo Getter
func (r AlibabaLstVendingShippingCallbackRequest) GetTradeFlowNo() string {
    return r.tradeFlowNo
}
// Code Setter
// 处理结果代码
func (r *AlibabaLstVendingShippingCallbackRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r AlibabaLstVendingShippingCallbackRequest) GetCode() string {
    return r.code
}
// Message Setter
// 处理结果代码描述
func (r *AlibabaLstVendingShippingCallbackRequest) SetMessage(message string) error {
    r.message = message
    r.Set("message", message)
    return nil
}

// Message Getter
func (r AlibabaLstVendingShippingCallbackRequest) GetMessage() string {
    return r.message
}
// Time Setter
// 出货时间
func (r *AlibabaLstVendingShippingCallbackRequest) SetTime(time string) error {
    r.time = time
    r.Set("time", time)
    return nil
}

// Time Getter
func (r AlibabaLstVendingShippingCallbackRequest) GetTime() string {
    return r.time
}
