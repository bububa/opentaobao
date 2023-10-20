package tmallsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSerivcecenterWorkcardInsuranceClaimAPIRequest 保险理赔回传工单记录 API请求
// tmall.serivcecenter.workcard.insurance.claim
//
// 保险理赔回传工单记录
type TmallSerivcecenterWorkcardInsuranceClaimAPIRequest struct {
	model.Params
	// 工单回传理赔信息入参
	_workcardInsuranceCallbackRequest *WorkcardInsuranceCallbackRequest
}

// NewTmallSerivcecenterWorkcardInsuranceClaimRequest 初始化TmallSerivcecenterWorkcardInsuranceClaimAPIRequest对象
func NewTmallSerivcecenterWorkcardInsuranceClaimRequest() *TmallSerivcecenterWorkcardInsuranceClaimAPIRequest {
	return &TmallSerivcecenterWorkcardInsuranceClaimAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallSerivcecenterWorkcardInsuranceClaimAPIRequest) Reset() {
	r._workcardInsuranceCallbackRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallSerivcecenterWorkcardInsuranceClaimAPIRequest) GetApiMethodName() string {
	return "tmall.serivcecenter.workcard.insurance.claim"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallSerivcecenterWorkcardInsuranceClaimAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallSerivcecenterWorkcardInsuranceClaimAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardInsuranceCallbackRequest is WorkcardInsuranceCallbackRequest Setter
// 工单回传理赔信息入参
func (r *TmallSerivcecenterWorkcardInsuranceClaimAPIRequest) SetWorkcardInsuranceCallbackRequest(_workcardInsuranceCallbackRequest *WorkcardInsuranceCallbackRequest) error {
	r._workcardInsuranceCallbackRequest = _workcardInsuranceCallbackRequest
	r.Set("workcard_insurance_callback_request", _workcardInsuranceCallbackRequest)
	return nil
}

// GetWorkcardInsuranceCallbackRequest WorkcardInsuranceCallbackRequest Getter
func (r TmallSerivcecenterWorkcardInsuranceClaimAPIRequest) GetWorkcardInsuranceCallbackRequest() *WorkcardInsuranceCallbackRequest {
	return r._workcardInsuranceCallbackRequest
}

var poolTmallSerivcecenterWorkcardInsuranceClaimAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallSerivcecenterWorkcardInsuranceClaimRequest()
	},
}

// GetTmallSerivcecenterWorkcardInsuranceClaimRequest 从 sync.Pool 获取 TmallSerivcecenterWorkcardInsuranceClaimAPIRequest
func GetTmallSerivcecenterWorkcardInsuranceClaimAPIRequest() *TmallSerivcecenterWorkcardInsuranceClaimAPIRequest {
	return poolTmallSerivcecenterWorkcardInsuranceClaimAPIRequest.Get().(*TmallSerivcecenterWorkcardInsuranceClaimAPIRequest)
}

// ReleaseTmallSerivcecenterWorkcardInsuranceClaimAPIRequest 将 TmallSerivcecenterWorkcardInsuranceClaimAPIRequest 放入 sync.Pool
func ReleaseTmallSerivcecenterWorkcardInsuranceClaimAPIRequest(v *TmallSerivcecenterWorkcardInsuranceClaimAPIRequest) {
	v.Reset()
	poolTmallSerivcecenterWorkcardInsuranceClaimAPIRequest.Put(v)
}
