package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代扣详情查询 APIRequest
alibaba.idle.agreement.pay.query

查询代扣结果
*/
type AlibabaIdleAgreementPayQueryRequest struct {
    model.Params

    // 入参
    param   *AgreementPayBillQueryParam 

}

func NewAlibabaIdleAgreementPayQueryRequest() *AlibabaIdleAgreementPayQueryRequest{
    return &AlibabaIdleAgreementPayQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleAgreementPayQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.agreement.pay.query"
}

func (r AlibabaIdleAgreementPayQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleAgreementPayQueryRequest) SetParam(param *AgreementPayBillQueryParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaIdleAgreementPayQueryRequest) GetParam() *AgreementPayBillQueryParam {
    return r.param
}

