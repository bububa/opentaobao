package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardVerifycodeResendAPIRequest 重发核销码 API请求
// tmall.servicecenter.workcard.verifycode.resend
//
// 重发核销码
type TmallServicecenterWorkcardVerifycodeResendAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
	// 门店/网点id
	_serviceStoreId int64
}

// NewTmallServicecenterWorkcardVerifycodeResendRequest 初始化TmallServicecenterWorkcardVerifycodeResendAPIRequest对象
func NewTmallServicecenterWorkcardVerifycodeResendRequest() *TmallServicecenterWorkcardVerifycodeResendAPIRequest {
	return &TmallServicecenterWorkcardVerifycodeResendAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkcardVerifycodeResendAPIRequest) Reset() {
	r._workcardId = 0
	r._serviceStoreId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardVerifycodeResendAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.verifycode.resend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardVerifycodeResendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardVerifycodeResendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardVerifycodeResendAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallServicecenterWorkcardVerifycodeResendAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetServiceStoreId is ServiceStoreId Setter
// 门店/网点id
func (r *TmallServicecenterWorkcardVerifycodeResendAPIRequest) SetServiceStoreId(_serviceStoreId int64) error {
	r._serviceStoreId = _serviceStoreId
	r.Set("service_store_id", _serviceStoreId)
	return nil
}

// GetServiceStoreId ServiceStoreId Getter
func (r TmallServicecenterWorkcardVerifycodeResendAPIRequest) GetServiceStoreId() int64 {
	return r._serviceStoreId
}

var poolTmallServicecenterWorkcardVerifycodeResendAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkcardVerifycodeResendRequest()
	},
}

// GetTmallServicecenterWorkcardVerifycodeResendRequest 从 sync.Pool 获取 TmallServicecenterWorkcardVerifycodeResendAPIRequest
func GetTmallServicecenterWorkcardVerifycodeResendAPIRequest() *TmallServicecenterWorkcardVerifycodeResendAPIRequest {
	return poolTmallServicecenterWorkcardVerifycodeResendAPIRequest.Get().(*TmallServicecenterWorkcardVerifycodeResendAPIRequest)
}

// ReleaseTmallServicecenterWorkcardVerifycodeResendAPIRequest 将 TmallServicecenterWorkcardVerifycodeResendAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkcardVerifycodeResendAPIRequest(v *TmallServicecenterWorkcardVerifycodeResendAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkcardVerifycodeResendAPIRequest.Put(v)
}
