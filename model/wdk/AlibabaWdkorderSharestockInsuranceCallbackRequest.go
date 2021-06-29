package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
共享库存订单投保后回传保单号 API请求
alibaba.wdkorder.sharestock.insurance.callback

共享库存订单投保消息获取
*/
type AlibabaWdkorderSharestockInsuranceCallbackRequest struct {
    model.Params
    // 淘宝交易子订单ID
    tbSubOrderId   int64
    // 投保单ID
    insuranceId   string
}

// 初始化AlibabaWdkorderSharestockInsuranceCallbackRequest对象
func NewAlibabaWdkorderSharestockInsuranceCallbackRequest() *AlibabaWdkorderSharestockInsuranceCallbackRequest{
    return &AlibabaWdkorderSharestockInsuranceCallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockInsuranceCallbackRequest) GetApiMethodName() string {
    return "alibaba.wdkorder.sharestock.insurance.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockInsuranceCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TbSubOrderId Setter
// 淘宝交易子订单ID
func (r *AlibabaWdkorderSharestockInsuranceCallbackRequest) SetTbSubOrderId(tbSubOrderId int64) error {
    r.tbSubOrderId = tbSubOrderId
    r.Set("tb_sub_order_id", tbSubOrderId)
    return nil
}

// TbSubOrderId Getter
func (r AlibabaWdkorderSharestockInsuranceCallbackRequest) GetTbSubOrderId() int64 {
    return r.tbSubOrderId
}
// InsuranceId Setter
// 投保单ID
func (r *AlibabaWdkorderSharestockInsuranceCallbackRequest) SetInsuranceId(insuranceId string) error {
    r.insuranceId = insuranceId
    r.Set("insurance_id", insuranceId)
    return nil
}

// InsuranceId Getter
func (r AlibabaWdkorderSharestockInsuranceCallbackRequest) GetInsuranceId() string {
    return r.insuranceId
}
