package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleAgreementPayQueryAPIRequest 代扣详情查询 API请求
// alibaba.idle.agreement.pay.query
//
// 查询代扣结果
type AlibabaIdleAgreementPayQueryAPIRequest struct {
	model.Params
	// 入参
	_param *AgreementPayBillQueryParam
}

// NewAlibabaIdleAgreementPayQueryRequest 初始化AlibabaIdleAgreementPayQueryAPIRequest对象
func NewAlibabaIdleAgreementPayQueryRequest() *AlibabaIdleAgreementPayQueryAPIRequest {
	return &AlibabaIdleAgreementPayQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleAgreementPayQueryAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleAgreementPayQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.agreement.pay.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleAgreementPayQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleAgreementPayQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaIdleAgreementPayQueryAPIRequest) SetParam(_param *AgreementPayBillQueryParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaIdleAgreementPayQueryAPIRequest) GetParam() *AgreementPayBillQueryParam {
	return r._param
}

var poolAlibabaIdleAgreementPayQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleAgreementPayQueryRequest()
	},
}

// GetAlibabaIdleAgreementPayQueryRequest 从 sync.Pool 获取 AlibabaIdleAgreementPayQueryAPIRequest
func GetAlibabaIdleAgreementPayQueryAPIRequest() *AlibabaIdleAgreementPayQueryAPIRequest {
	return poolAlibabaIdleAgreementPayQueryAPIRequest.Get().(*AlibabaIdleAgreementPayQueryAPIRequest)
}

// ReleaseAlibabaIdleAgreementPayQueryAPIRequest 将 AlibabaIdleAgreementPayQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleAgreementPayQueryAPIRequest(v *AlibabaIdleAgreementPayQueryAPIRequest) {
	v.Reset()
	poolAlibabaIdleAgreementPayQueryAPIRequest.Put(v)
}
