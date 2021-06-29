package cainiaoecc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟控制塔包裹滞留异常统计信息获取 APIRequest
cainiao.ecc.exceptions.delay.count

菜鸟控制塔包裹滞留异常统计信息获取
*/
type CainiaoEccExceptionsDelayCountRequest struct {
    model.Params

}

func NewCainiaoEccExceptionsDelayCountRequest() *CainiaoEccExceptionsDelayCountRequest{
    return &CainiaoEccExceptionsDelayCountRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoEccExceptionsDelayCountRequest) GetApiMethodName() string {
    return "cainiao.ecc.exceptions.delay.count"
}

func (r CainiaoEccExceptionsDelayCountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


