package interactvip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员淘气值获取 APIRequest
alibaba.interact.vip.get

提供用户淘气值&用户角色身份查询
*/
type AlibabaInteractVipGetRequest struct {
    model.Params

}

func NewAlibabaInteractVipGetRequest() *AlibabaInteractVipGetRequest{
    return &AlibabaInteractVipGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractVipGetRequest) GetApiMethodName() string {
    return "alibaba.interact.vip.get"
}

func (r AlibabaInteractVipGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


