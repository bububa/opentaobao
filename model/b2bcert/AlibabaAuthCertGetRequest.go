package b2bcert

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取证书数据 API请求
alibaba.auth.cert.get

获取证书数据
*/
type AlibabaAuthCertGetRequest struct {
    model.Params
    // 认证商
    _provider   string
    // 证书数据
    _receiveInfo   string
}

// 初始化AlibabaAuthCertGetRequest对象
func NewAlibabaAuthCertGetRequest() *AlibabaAuthCertGetRequest{
    return &AlibabaAuthCertGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAuthCertGetRequest) GetApiMethodName() string {
    return "alibaba.auth.cert.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAuthCertGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Provider Setter
// 认证商
func (r *AlibabaAuthCertGetRequest) SetProvider(_provider string) error {
    r._provider = _provider
    r.Set("provider", _provider)
    return nil
}

// Provider Getter
func (r AlibabaAuthCertGetRequest) GetProvider() string {
    return r._provider
}
// ReceiveInfo Setter
// 证书数据
func (r *AlibabaAuthCertGetRequest) SetReceiveInfo(_receiveInfo string) error {
    r._receiveInfo = _receiveInfo
    r.Set("receive_info", _receiveInfo)
    return nil
}

// ReceiveInfo Getter
func (r AlibabaAuthCertGetRequest) GetReceiveInfo() string {
    return r._receiveInfo
}
