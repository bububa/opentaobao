package cainiaoncwl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
农村物流集货单查询接口 APIRequest
cainiao.ncwl.jhorder.query

提供给接入商家，查询农村物流集货单
*/
type CainiaoNcwlJhorderQueryRequest struct {
    model.Params

    // 1
    param0   *JhRequest 

}

func NewCainiaoNcwlJhorderQueryRequest() *CainiaoNcwlJhorderQueryRequest{
    return &CainiaoNcwlJhorderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoNcwlJhorderQueryRequest) GetApiMethodName() string {
    return "cainiao.ncwl.jhorder.query"
}

func (r CainiaoNcwlJhorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoNcwlJhorderQueryRequest) SetParam0(param0 *JhRequest) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r CainiaoNcwlJhorderQueryRequest) GetParam0() *JhRequest {
    return r.param0
}

