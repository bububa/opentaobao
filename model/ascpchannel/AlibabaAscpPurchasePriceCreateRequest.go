package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ascp采购价写入接口 APIRequest
alibaba.ascp.purchase.price.create

供应链平台采购价创建或修改接口
*/
type AlibabaAscpPurchasePriceCreateRequest struct {
    model.Params

    // 采购价创建/更新请求
    createRequest   *Request 

}

func NewAlibabaAscpPurchasePriceCreateRequest() *AlibabaAscpPurchasePriceCreateRequest{
    return &AlibabaAscpPurchasePriceCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpPurchasePriceCreateRequest) GetApiMethodName() string {
    return "alibaba.ascp.purchase.price.create"
}

func (r AlibabaAscpPurchasePriceCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpPurchasePriceCreateRequest) SetCreateRequest(createRequest *Request) error {
    r.createRequest = createRequest
    r.Set("create_request", createRequest)
    return nil
}

func (r AlibabaAscpPurchasePriceCreateRequest) GetCreateRequest() *Request {
    return r.createRequest
}

