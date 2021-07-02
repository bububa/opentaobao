package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleAgreementPayAPIRequest 闲鱼平台商户代扣 API请求
// alibaba.idle.agreement.pay
//
// 闲鱼平台代扣能力：用户和闲鱼签约代扣协议 服务商通过直付通产品挂载成为闲鱼二级商户 来完成用户和服务商的结算
type AlibabaIdleAgreementPayAPIRequest struct {
	model.Params
	// 协议代扣参数
	_agreementPayParam *AgreementPayParam
}

// NewAlibabaIdleAgreementPayRequest 初始化AlibabaIdleAgreementPayAPIRequest对象
func NewAlibabaIdleAgreementPayRequest() *AlibabaIdleAgreementPayAPIRequest {
	return &AlibabaIdleAgreementPayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleAgreementPayAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.agreement.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleAgreementPayAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAgreementPayParam is AgreementPayParam Setter
// 协议代扣参数
func (r *AlibabaIdleAgreementPayAPIRequest) SetAgreementPayParam(_agreementPayParam *AgreementPayParam) error {
	r._agreementPayParam = _agreementPayParam
	r.Set("agreement_pay_param", _agreementPayParam)
	return nil
}

// GetAgreementPayParam AgreementPayParam Getter
func (r AlibabaIdleAgreementPayAPIRequest) GetAgreementPayParam() *AgreementPayParam {
	return r._agreementPayParam
}
