package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代扣详情查询 API请求
alibaba.idle.agreement.pay.query

查询代扣结果
*/
type AlibabaIdleAgreementPayQueryRequest struct {
    model.Params
    // 入参
    param   *AgreementPayBillQueryParam
}

// 初始化AlibabaIdleAgreementPayQueryRequest对象
func NewAlibabaIdleAgreementPayQueryRequest() *AlibabaIdleAgreementPayQueryRequest{
    return &AlibabaIdleAgreementPayQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleAgreementPayQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.agreement.pay.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleAgreementPayQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaIdleAgreementPayQueryRequest) SetParam(param *AgreementPayBillQueryParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaIdleAgreementPayQueryRequest) GetParam() *AgreementPayBillQueryParam {
    return r.param
}
