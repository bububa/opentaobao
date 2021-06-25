package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取天猫互动奖池列表 APIRequest
alibaba.interact.isvadmin.allponds

获取天猫互动奖池列表
*/
type AlibabaInteractIsvadminAllpondsRequest struct {
    model.Params

}

func NewAlibabaInteractIsvadminAllpondsRequest() *AlibabaInteractIsvadminAllpondsRequest{
    return &AlibabaInteractIsvadminAllpondsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractIsvadminAllpondsRequest) GetApiMethodName() string {
    return "alibaba.interact.isvadmin.allponds"
}

func (r AlibabaInteractIsvadminAllpondsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


