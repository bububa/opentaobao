package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工人信息查询 API请求
tmall.servicecenter.worker.query

查询服务商对应的工人信息
*/
type TmallServicecenterWorkerQueryRequest struct {
    model.Params
    // 身份证号
    identityId   string
}

// 初始化TmallServicecenterWorkerQueryRequest对象
func NewTmallServicecenterWorkerQueryRequest() *TmallServicecenterWorkerQueryRequest{
    return &TmallServicecenterWorkerQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerQueryRequest) GetApiMethodName() string {
    return "tmall.servicecenter.worker.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IdentityId Setter
// 身份证号
func (r *TmallServicecenterWorkerQueryRequest) SetIdentityId(identityId string) error {
    r.identityId = identityId
    r.Set("identity_id", identityId)
    return nil
}

// IdentityId Getter
func (r TmallServicecenterWorkerQueryRequest) GetIdentityId() string {
    return r.identityId
}
