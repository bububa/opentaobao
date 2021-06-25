package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工人信息查询 APIRequest
tmall.servicecenter.worker.query

查询服务商对应的工人信息
*/
type TmallServicecenterWorkerQueryRequest struct {
    model.Params

    // 身份证号
    identityId   string 

}

func NewTmallServicecenterWorkerQueryRequest() *TmallServicecenterWorkerQueryRequest{
    return &TmallServicecenterWorkerQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkerQueryRequest) GetApiMethodName() string {
    return "tmall.servicecenter.worker.query"
}

func (r TmallServicecenterWorkerQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkerQueryRequest) SetIdentityId(identityId string) error {
    r.identityId = identityId
    r.Set("identity_id", identityId)
    return nil
}

func (r TmallServicecenterWorkerQueryRequest) GetIdentityId() string {
    return r.identityId
}

