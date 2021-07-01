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
type AlibabaAlscCrmMarketingEncryptAPIRequest struct {
    model.Params
    // 参数
    _param   string
}

// 初始化AlibabaAlscCrmMarketingEncryptAPIRequest对象
func NewAlibabaAlscCrmMarketingEncryptRequest() *AlibabaAlscCrmMarketingEncryptAPIRequest{
    return &AlibabaAlscCrmMarketingEncryptAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmMarketingEncryptAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.marketing.encrypt"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmMarketingEncryptAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 参数
func (r *AlibabaAlscCrmMarketingEncryptAPIRequest) SetParam(_param string) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaAlscCrmMarketingEncryptAPIRequest) GetParam() string {
    return r._param
}
