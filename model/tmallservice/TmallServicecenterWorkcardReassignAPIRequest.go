package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardreassignAPIRequest 工单改派门店 API请求
// tmall.servicecenter.workcard.reassign
//
// 工单改派门店
type TmallservicecenterworkcardreassignAPIRequest struct {
	model.Params
	// 请求入参
	_reassignStoreRequest *ReassignStoreRequest
}

// NewTmallservicecenterworkcardreassignRequest 初始化TmallservicecenterworkcardreassignAPIRequest对象
func NewTmallservicecenterworkcardreassignRequest() *TmallservicecenterworkcardreassignAPIRequest {
	return &TmallservicecenterworkcardreassignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardreassignAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.reassign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardreassignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardreassignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReassignStoreRequest is ReassignStoreRequest Setter
// 请求入参
func (r *TmallservicecenterworkcardreassignAPIRequest) SetReassignStoreRequest(_reassignStoreRequest *ReassignStoreRequest) error {
	r._reassignStoreRequest = _reassignStoreRequest
	r.Set("reassign_store_request", _reassignStoreRequest)
	return nil
}

// GetReassignStoreRequest ReassignStoreRequest Getter
func (r TmallservicecenterworkcardreassignAPIRequest) GetReassignStoreRequest() *ReassignStoreRequest {
	return r._reassignStoreRequest
}
