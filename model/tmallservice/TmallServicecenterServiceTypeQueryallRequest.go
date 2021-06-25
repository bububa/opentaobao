package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务供应链服务类型 APIRequest
tmall.servicecenter.service.type.queryall

查询天猫服务类型列表
*/
type TmallServicecenterServiceTypeQueryallRequest struct {
    model.Params

}

func NewTmallServicecenterServiceTypeQueryallRequest() *TmallServicecenterServiceTypeQueryallRequest{
    return &TmallServicecenterServiceTypeQueryallRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterServiceTypeQueryallRequest) GetApiMethodName() string {
    return "tmall.servicecenter.service.type.queryall"
}

func (r TmallServicecenterServiceTypeQueryallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


