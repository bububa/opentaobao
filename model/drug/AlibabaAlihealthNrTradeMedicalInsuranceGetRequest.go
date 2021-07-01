package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康医保支付信息获取 API请求
alibaba.alihealth.nr.trade.medical.insurance.get

阿里健康医保支付信息获取
*/
type AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest struct {
    model.Params
    // 淘宝订单ID
    _orderId   int64
}

// 初始化AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest对象
func NewAlibabaAlihealthNrTradeMedicalInsuranceGetRequest() *AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest{
    return &AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.nr.trade.medical.insurance.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 淘宝订单ID
func (r *AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest) GetOrderId() int64 {
    return r._orderId
}
