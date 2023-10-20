package idle

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleAgreementPayAPIRequest) Reset() {
	r._agreementPayParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleAgreementPayAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.agreement.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleAgreementPayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleAgreementPayAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaIdleAgreementPayAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleAgreementPayRequest()
	},
}

// GetAlibabaIdleAgreementPayRequest 从 sync.Pool 获取 AlibabaIdleAgreementPayAPIRequest
func GetAlibabaIdleAgreementPayAPIRequest() *AlibabaIdleAgreementPayAPIRequest {
	return poolAlibabaIdleAgreementPayAPIRequest.Get().(*AlibabaIdleAgreementPayAPIRequest)
}

// ReleaseAlibabaIdleAgreementPayAPIRequest 将 AlibabaIdleAgreementPayAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleAgreementPayAPIRequest(v *AlibabaIdleAgreementPayAPIRequest) {
	v.Reset()
	poolAlibabaIdleAgreementPayAPIRequest.Put(v)
}
