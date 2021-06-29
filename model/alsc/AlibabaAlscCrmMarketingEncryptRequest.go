package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
加密 API请求
alibaba.alsc.crm.marketing.encrypt

加密
*/
type AlibabaAlscCrmMarketingEncryptRequest struct {
    model.Params
    // 参数
    _param   string
}

// 初始化AlibabaAlscCrmMarketingEncryptRequest对象
func NewAlibabaAlscCrmMarketingEncryptRequest() *AlibabaAlscCrmMarketingEncryptRequest{
    return &AlibabaAlscCrmMarketingEncryptRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmMarketingEncryptRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.marketing.encrypt"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmMarketingEncryptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 参数
func (r *AlibabaAlscCrmMarketingEncryptRequest) SetParam(_param string) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaAlscCrmMarketingEncryptRequest) GetParam() string {
    return r._param
}
