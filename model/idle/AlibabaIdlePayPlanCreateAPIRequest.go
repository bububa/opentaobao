package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdlePayPlanCreateAPIRequest 创建代扣计划 API请求
// alibaba.idle.pay.plan.create
//
// 闲鱼平台代扣能力：
// 1、用户和闲鱼签约代扣协议 服务商通过直付通产品挂载成为闲鱼二级商户 来完成用户和服务商的结算
// 2、创建代扣计划
type AlibabaIdlePayPlanCreateAPIRequest struct {
	model.Params
	// 业务入参
	_agreementPayPlanParam *AgreementPayPlanParam
}

// NewAlibabaIdlePayPlanCreateRequest 初始化AlibabaIdlePayPlanCreateAPIRequest对象
func NewAlibabaIdlePayPlanCreateRequest() *AlibabaIdlePayPlanCreateAPIRequest {
	return &AlibabaIdlePayPlanCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdlePayPlanCreateAPIRequest) Reset() {
	r._agreementPayPlanParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdlePayPlanCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.pay.plan.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdlePayPlanCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdlePayPlanCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgreementPayPlanParam is AgreementPayPlanParam Setter
// 业务入参
func (r *AlibabaIdlePayPlanCreateAPIRequest) SetAgreementPayPlanParam(_agreementPayPlanParam *AgreementPayPlanParam) error {
	r._agreementPayPlanParam = _agreementPayPlanParam
	r.Set("agreement_pay_plan_param", _agreementPayPlanParam)
	return nil
}

// GetAgreementPayPlanParam AgreementPayPlanParam Getter
func (r AlibabaIdlePayPlanCreateAPIRequest) GetAgreementPayPlanParam() *AgreementPayPlanParam {
	return r._agreementPayPlanParam
}

var poolAlibabaIdlePayPlanCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdlePayPlanCreateRequest()
	},
}

// GetAlibabaIdlePayPlanCreateRequest 从 sync.Pool 获取 AlibabaIdlePayPlanCreateAPIRequest
func GetAlibabaIdlePayPlanCreateAPIRequest() *AlibabaIdlePayPlanCreateAPIRequest {
	return poolAlibabaIdlePayPlanCreateAPIRequest.Get().(*AlibabaIdlePayPlanCreateAPIRequest)
}

// ReleaseAlibabaIdlePayPlanCreateAPIRequest 将 AlibabaIdlePayPlanCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdlePayPlanCreateAPIRequest(v *AlibabaIdlePayPlanCreateAPIRequest) {
	v.Reset()
	poolAlibabaIdlePayPlanCreateAPIRequest.Put(v)
}
