package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
本地生活风控数据接入 API请求
alibaba.alsc.kms.access

第三方使用本地生活数据对外提供服务，上报访问日志信息接口
*/
type AlibabaAlscKmsAccessRequest struct {
    model.Params
    // 接入参数
    requestdata   string
}

// 初始化AlibabaAlscKmsAccessRequest对象
func NewAlibabaAlscKmsAccessRequest() *AlibabaAlscKmsAccessRequest{
    return &AlibabaAlscKmsAccessRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscKmsAccessRequest) GetApiMethodName() string {
    return "alibaba.alsc.kms.access"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscKmsAccessRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Requestdata Setter
// 接入参数
func (r *AlibabaAlscKmsAccessRequest) SetRequestdata(requestdata string) error {
    r.requestdata = requestdata
    r.Set("requestdata", requestdata)
    return nil
}

// Requestdata Getter
func (r AlibabaAlscKmsAccessRequest) GetRequestdata() string {
    return r.requestdata
}
