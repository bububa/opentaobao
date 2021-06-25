package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品 APIRequest
alibaba.data.item.get

获取商品信息，作为客户端Weex鉴权的虚拟api
*/
type AlibabaDataItemGetRequest struct {
    model.Params

    // 获取商品信息，作为客户端Weex鉴权的虚拟api
    unNamed   string 

}

func NewAlibabaDataItemGetRequest() *AlibabaDataItemGetRequest{
    return &AlibabaDataItemGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDataItemGetRequest) GetApiMethodName() string {
    return "alibaba.data.item.get"
}

func (r AlibabaDataItemGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDataItemGetRequest) SetUnNamed(unNamed string) error {
    r.unNamed = unNamed
    r.Set("un_named", unNamed)
    return nil
}

func (r AlibabaDataItemGetRequest) GetUnNamed() string {
    return r.unNamed
}

