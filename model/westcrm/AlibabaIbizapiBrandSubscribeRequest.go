package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关注品牌号 APIRequest
alibaba.ibizapi.brand.subscribe

关注品牌号服务
*/
type AlibabaIbizapiBrandSubscribeRequest struct {
    model.Params

}

func NewAlibabaIbizapiBrandSubscribeRequest() *AlibabaIbizapiBrandSubscribeRequest{
    return &AlibabaIbizapiBrandSubscribeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIbizapiBrandSubscribeRequest) GetApiMethodName() string {
    return "alibaba.ibizapi.brand.subscribe"
}

func (r AlibabaIbizapiBrandSubscribeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


