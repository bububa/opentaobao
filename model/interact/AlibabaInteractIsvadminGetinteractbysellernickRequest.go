package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据sellerNick获取互动实例列表 APIRequest
alibaba.interact.isvadmin.getinteractbysellernick

根据sellerNick获取互动实例列表
*/
type AlibabaInteractIsvadminGetinteractbysellernickRequest struct {
    model.Params

}

func NewAlibabaInteractIsvadminGetinteractbysellernickRequest() *AlibabaInteractIsvadminGetinteractbysellernickRequest{
    return &AlibabaInteractIsvadminGetinteractbysellernickRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractIsvadminGetinteractbysellernickRequest) GetApiMethodName() string {
    return "alibaba.interact.isvadmin.getinteractbysellernick"
}

func (r AlibabaInteractIsvadminGetinteractbysellernickRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


