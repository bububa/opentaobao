package lstvending

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
售货机出货回传接口 APIRequest
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

func NewAlibabaLstVendingShippingCallbackRequest() *AlibabaLstVendingShippingCallbackRequest{
    return &AlibabaLstVendingShippingCallbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstVendingShippingCallbackRequest) GetApiMethodName() string {
    return "alibaba.lst.vending.shipping.callback"
}

func (r AlibabaLstVendingShippingCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstVendingShippingCallbackRequest) SetEquipmentCode(equipmentCode string) error {
    r.equipmentCode = equipmentCode
    r.Set("equipment_code", equipmentCode)
    return nil
}

func (r AlibabaLstVendingShippingCallbackRequest) GetEquipmentCode() string {
    return r.equipmentCode
}

func (r *AlibabaLstVendingShippingCallbackRequest) SetTradeFlowNo(tradeFlowNo string) error {
    r.tradeFlowNo = tradeFlowNo
    r.Set("trade_flow_no", tradeFlowNo)
    return nil
}

func (r AlibabaLstVendingShippingCallbackRequest) GetTradeFlowNo() string {
    return r.tradeFlowNo
}

func (r *AlibabaLstVendingShippingCallbackRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaLstVendingShippingCallbackRequest) GetCode() string {
    return r.code
}

func (r *AlibabaLstVendingShippingCallbackRequest) SetMessage(message string) error {
    r.message = message
    r.Set("message", message)
    return nil
}

func (r AlibabaLstVendingShippingCallbackRequest) GetMessage() string {
    return r.message
}

func (r *AlibabaLstVendingShippingCallbackRequest) SetTime(time string) error {
    r.time = time
    r.Set("time", time)
    return nil
}

func (r AlibabaLstVendingShippingCallbackRequest) GetTime() string {
    return r.time
}

