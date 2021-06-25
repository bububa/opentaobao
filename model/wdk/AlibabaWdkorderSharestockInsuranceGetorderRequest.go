package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
共享库存订单投保消息获取 APIRequest
alibaba.wdkorder.sharestock.insurance.getorder

共享库存订单投保消息获取
*/
type AlibabaWdkorderSharestockInsuranceGetorderRequest struct {
    model.Params

    // 淘宝子订单ID
    tbSubOrderId   int64 

}

func NewAlibabaWdkorderSharestockInsuranceGetorderRequest() *AlibabaWdkorderSharestockInsuranceGetorderRequest{
    return &AlibabaWdkorderSharestockInsuranceGetorderRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkorderSharestockInsuranceGetorderRequest) GetApiMethodName() string {
    return "alibaba.wdkorder.sharestock.insurance.getorder"
}

func (r AlibabaWdkorderSharestockInsuranceGetorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkorderSharestockInsuranceGetorderRequest) SetTbSubOrderId(tbSubOrderId int64) error {
    r.tbSubOrderId = tbSubOrderId
    r.Set("tb_sub_order_id", tbSubOrderId)
    return nil
}

func (r AlibabaWdkorderSharestockInsuranceGetorderRequest) GetTbSubOrderId() int64 {
    return r.tbSubOrderId
}

