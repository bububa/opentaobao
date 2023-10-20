package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleagreementpayAPIRequest 闲鱼平台商户代扣 API请求
// alibaba.idle.agreement.pay
//
// 闲鱼平台代扣能力：用户和闲鱼签约代扣协议 服务商通过直付通产品挂载成为闲鱼二级商户 来完成用户和服务商的结算
type AlibabaidleagreementpayAPIRequest struct {
	model.Params
	// 协议代扣参数
	_agreementPayParam *AgreementPayParam
}

// NewAlibabaidleagreementpayRequest 初始化AlibabaidleagreementpayAPIRequest对象
func NewAlibabaidleagreementpayRequest() *AlibabaidleagreementpayAPIRequest {
	return &AlibabaidleagreementpayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleagreementpayAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.agreement.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleagreementpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleagreementpayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgreementPayParam is AgreementPayParam Setter
// 协议代扣参数
func (r *AlibabaidleagreementpayAPIRequest) SetAgreementPayParam(_agreementPayParam *AgreementPayParam) error {
	r._agreementPayParam = _agreementPayParam
	r.Set("agreement_pay_param", _agreementPayParam)
	return nil
}

// GetAgreementPayParam AgreementPayParam Getter
func (r AlibabaidleagreementpayAPIRequest) GetAgreementPayParam() *AgreementPayParam {
	return r._agreementPayParam
}
