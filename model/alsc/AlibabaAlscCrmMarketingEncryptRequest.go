package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
加密 APIRequest
alibaba.alsc.crm.marketing.encrypt

加密
*/
type AlibabaAlscCrmMarketingEncryptRequest struct {
    model.Params

    // 参数
    param   string 

}

func NewAlibabaAlscCrmMarketingEncryptRequest() *AlibabaAlscCrmMarketingEncryptRequest{
    return &AlibabaAlscCrmMarketingEncryptRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmMarketingEncryptRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.marketing.encrypt"
}

func (r AlibabaAlscCrmMarketingEncryptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmMarketingEncryptRequest) SetParam(param string) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaAlscCrmMarketingEncryptRequest) GetParam() string {
    return r.param
}

