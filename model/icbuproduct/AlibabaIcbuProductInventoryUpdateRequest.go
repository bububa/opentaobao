package icbuproduct

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
icbu商品库存更新 APIRequest
alibaba.icbu.product.inventory.update

更新库存信息
*/
type AlibabaIcbuProductInventoryUpdateRequest struct {
    model.Params

    // 更新请求
    requestParam   *ProductInventoryRequest 

}

func NewAlibabaIcbuProductInventoryUpdateRequest() *AlibabaIcbuProductInventoryUpdateRequest{
    return &AlibabaIcbuProductInventoryUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuProductInventoryUpdateRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.inventory.update"
}

func (r AlibabaIcbuProductInventoryUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuProductInventoryUpdateRequest) SetRequestParam(requestParam *ProductInventoryRequest) error {
    r.requestParam = requestParam
    r.Set("request_param", requestParam)
    return nil
}

func (r AlibabaIcbuProductInventoryUpdateRequest) GetRequestParam() *ProductInventoryRequest {
    return r.requestParam
}

