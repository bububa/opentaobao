package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleAgreementPayQueryAPIRequest
代扣详情查询 API请求
alibaba.idle.agreement.pay.query

查询代扣结果 */
type AlibabaIdleAgreementPayQueryAPIRequest struct {
	model.Params
	// 入参
	_param *AgreementPayBillQueryParam
}

// NewAlibabaIdleAgreementPayQueryRequest 初始化AlibabaIdleAgreementPayQueryAPIRequest对象
func NewAlibabaIdleAgreementPayQueryRequest() *AlibabaIdleAgreementPayQueryAPIRequest {
	return &AlibabaIdleAgreementPayQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleAgreementPayQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.agreement.pay.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleAgreementPayQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 入参
func (r *AlibabaIdleAgreementPayQueryAPIRequest) SetParam(_param *AgreementPayBillQueryParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaIdleAgreementPayQueryAPIRequest) GetParam() *AgreementPayBillQueryParam {
	return r._param
}
