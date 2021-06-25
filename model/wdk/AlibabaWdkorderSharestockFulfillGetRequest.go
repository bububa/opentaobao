package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商户订单履约数据获取 APIRequest
alibaba.wdkorder.sharestock.fulfill.get

商户订单履约数据获取
*/
type AlibabaWdkorderSharestockFulfillGetRequest struct {
    model.Params

    // 履约单ID
    fulfillOrderId   string 

}

func NewAlibabaWdkorderSharestockFulfillGetRequest() *AlibabaWdkorderSharestockFulfillGetRequest{
    return &AlibabaWdkorderSharestockFulfillGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkorderSharestockFulfillGetRequest) GetApiMethodName() string {
    return "alibaba.wdkorder.sharestock.fulfill.get"
}

func (r AlibabaWdkorderSharestockFulfillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkorderSharestockFulfillGetRequest) SetFulfillOrderId(fulfillOrderId string) error {
    r.fulfillOrderId = fulfillOrderId
    r.Set("fulfill_order_id", fulfillOrderId)
    return nil
}

func (r AlibabaWdkorderSharestockFulfillGetRequest) GetFulfillOrderId() string {
    return r.fulfillOrderId
}

