package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardReassignAPIRequest 工单改派门店 API请求
// tmall.servicecenter.workcard.reassign
//
// 工单改派门店
type TmallServicecenterWorkcardReassignAPIRequest struct {
	model.Params
	// 请求入参
	_reassignStoreRequest *ReassignStoreRequest
}

// NewTmallServicecenterWorkcardReassignRequest 初始化TmallServicecenterWorkcardReassignAPIRequest对象
func NewTmallServicecenterWorkcardReassignRequest() *TmallServicecenterWorkcardReassignAPIRequest {
	return &TmallServicecenterWorkcardReassignAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkcardReassignAPIRequest) Reset() {
	r._reassignStoreRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardReassignAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.reassign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardReassignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardReassignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReassignStoreRequest is ReassignStoreRequest Setter
// 请求入参
func (r *TmallServicecenterWorkcardReassignAPIRequest) SetReassignStoreRequest(_reassignStoreRequest *ReassignStoreRequest) error {
	r._reassignStoreRequest = _reassignStoreRequest
	r.Set("reassign_store_request", _reassignStoreRequest)
	return nil
}

// GetReassignStoreRequest ReassignStoreRequest Getter
func (r TmallServicecenterWorkcardReassignAPIRequest) GetReassignStoreRequest() *ReassignStoreRequest {
	return r._reassignStoreRequest
}

var poolTmallServicecenterWorkcardReassignAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkcardReassignRequest()
	},
}

// GetTmallServicecenterWorkcardReassignRequest 从 sync.Pool 获取 TmallServicecenterWorkcardReassignAPIRequest
func GetTmallServicecenterWorkcardReassignAPIRequest() *TmallServicecenterWorkcardReassignAPIRequest {
	return poolTmallServicecenterWorkcardReassignAPIRequest.Get().(*TmallServicecenterWorkcardReassignAPIRequest)
}

// ReleaseTmallServicecenterWorkcardReassignAPIRequest 将 TmallServicecenterWorkcardReassignAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkcardReassignAPIRequest(v *TmallServicecenterWorkcardReassignAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkcardReassignAPIRequest.Put(v)
}
