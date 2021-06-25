package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康医保支付信息获取 APIRequest
alibaba.alihealth.nr.trade.medical.insurance.get

阿里健康医保支付信息获取
*/
type AlibabaAlihealthNrTradeMedicalInsuranceGetRequest struct {
    model.Params

    // 淘宝订单ID
    orderId   int64 

}

func NewAlibabaAlihealthNrTradeMedicalInsuranceGetRequest() *AlibabaAlihealthNrTradeMedicalInsuranceGetRequest{
    return &AlibabaAlihealthNrTradeMedicalInsuranceGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthNrTradeMedicalInsuranceGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.nr.trade.medical.insurance.get"
}

func (r AlibabaAlihealthNrTradeMedicalInsuranceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthNrTradeMedicalInsuranceGetRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaAlihealthNrTradeMedicalInsuranceGetRequest) GetOrderId() int64 {
    return r.orderId
}

