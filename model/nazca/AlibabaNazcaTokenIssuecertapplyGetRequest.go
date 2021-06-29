package nazca

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据token获取出证申请信息 API请求
alibaba.nazca.token.issuecertapply.get

根据token获取出证申请信息
*/
type AlibabaNazcaTokenIssuecertapplyGetRequest struct {
    model.Params
    // token
    _token   string
}

// 初始化AlibabaNazcaTokenIssuecertapplyGetRequest对象
func NewAlibabaNazcaTokenIssuecertapplyGetRequest() *AlibabaNazcaTokenIssuecertapplyGetRequest{
    return &AlibabaNazcaTokenIssuecertapplyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNazcaTokenIssuecertapplyGetRequest) GetApiMethodName() string {
    return "alibaba.nazca.token.issuecertapply.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNazcaTokenIssuecertapplyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Token Setter
// token
func (r *AlibabaNazcaTokenIssuecertapplyGetRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaNazcaTokenIssuecertapplyGetRequest) GetToken() string {
    return r._token
}
