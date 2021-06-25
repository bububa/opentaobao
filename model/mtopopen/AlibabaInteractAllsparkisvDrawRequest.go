package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
allspark提供抽奖tida接口对应鉴权接口 APIRequest
alibaba.interact.allsparkisv.draw

该接口没有实际对外使用。只是内部鉴权使用，不会有三方应用调用
*/
type AlibabaInteractAllsparkisvDrawRequest struct {
    model.Params

    // ddd
    test   string 

    // dd
    ddd   string 

}

func NewAlibabaInteractAllsparkisvDrawRequest() *AlibabaInteractAllsparkisvDrawRequest{
    return &AlibabaInteractAllsparkisvDrawRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractAllsparkisvDrawRequest) GetApiMethodName() string {
    return "alibaba.interact.allsparkisv.draw"
}

func (r AlibabaInteractAllsparkisvDrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractAllsparkisvDrawRequest) SetTest(test string) error {
    r.test = test
    r.Set("test", test)
    return nil
}

func (r AlibabaInteractAllsparkisvDrawRequest) GetTest() string {
    return r.test
}

func (r *AlibabaInteractAllsparkisvDrawRequest) SetDdd(ddd string) error {
    r.ddd = ddd
    r.Set("ddd", ddd)
    return nil
}

func (r AlibabaInteractAllsparkisvDrawRequest) GetDdd() string {
    return r.ddd
}

