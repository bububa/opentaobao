package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店查询接口 APIRequest
alibaba.wdk.shop.query

根据门店code查询门店信息
*/
type AlibabaWdkShopQueryRequest struct {
    model.Params

    // 如果不传，返回所有
    ouCode   string 

}

func NewAlibabaWdkShopQueryRequest() *AlibabaWdkShopQueryRequest{
    return &AlibabaWdkShopQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkShopQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.shop.query"
}

func (r AlibabaWdkShopQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkShopQueryRequest) SetOuCode(ouCode string) error {
    r.ouCode = ouCode
    r.Set("ou_code", ouCode)
    return nil
}

func (r AlibabaWdkShopQueryRequest) GetOuCode() string {
    return r.ouCode
}

