package dmp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询人群服务 APIRequest
taobao.dmp.crowds.get

查询人群服务
*/
type TaobaoDmpCrowdsGetRequest struct {
    model.Params

}

func NewTaobaoDmpCrowdsGetRequest() *TaobaoDmpCrowdsGetRequest{
    return &TaobaoDmpCrowdsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDmpCrowdsGetRequest) GetApiMethodName() string {
    return "taobao.dmp.crowds.get"
}

func (r TaobaoDmpCrowdsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


