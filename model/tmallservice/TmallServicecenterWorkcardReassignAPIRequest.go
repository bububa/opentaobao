package tmallservice

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardReassignAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.reassign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardReassignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ReassignStoreRequest Setter
// 请求入参
func (r *TmallServicecenterWorkcardReassignAPIRequest) SetReassignStoreRequest(_reassignStoreRequest *ReassignStoreRequest) error {
	r._reassignStoreRequest = _reassignStoreRequest
	r.Set("reassign_store_request", _reassignStoreRequest)
	return nil
}

// Get ReassignStoreRequest Getter
func (r TmallServicecenterWorkcardReassignAPIRequest) GetReassignStoreRequest() *ReassignStoreRequest {
	return r._reassignStoreRequest
}
