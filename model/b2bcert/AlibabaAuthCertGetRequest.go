package b2bcert

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取证书数据 APIRequest
alibaba.auth.cert.get

获取证书数据
*/
type AlibabaAuthCertGetRequest struct {
    model.Params

    // 认证商
    provider   string 

    // 证书数据
    receiveInfo   string 

}

func NewAlibabaAuthCertGetRequest() *AlibabaAuthCertGetRequest{
    return &AlibabaAuthCertGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAuthCertGetRequest) GetApiMethodName() string {
    return "alibaba.auth.cert.get"
}

func (r AlibabaAuthCertGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAuthCertGetRequest) SetProvider(provider string) error {
    r.provider = provider
    r.Set("provider", provider)
    return nil
}

func (r AlibabaAuthCertGetRequest) GetProvider() string {
    return r.provider
}

func (r *AlibabaAuthCertGetRequest) SetReceiveInfo(receiveInfo string) error {
    r.receiveInfo = receiveInfo
    r.Set("receive_info", receiveInfo)
    return nil
}

func (r AlibabaAuthCertGetRequest) GetReceiveInfo() string {
    return r.receiveInfo
}

