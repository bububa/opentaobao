package tmallsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallSparepartsDetailsCreateAPIRequest 天猫蚁巢同步工单申请备件明细 API请求
// alibaba.tmall.spareparts.details.create
//
// 天猫蚁巢同步工单申请备件明细
type AlibabaTmallSparepartsDetailsCreateAPIRequest struct {
	model.Params
	// 请求入参
	_sparePartsDetailsSaveRequest *SparePartsDetailsSaveRequest
}

// NewAlibabaTmallSparepartsDetailsCreateRequest 初始化AlibabaTmallSparepartsDetailsCreateAPIRequest对象
func NewAlibabaTmallSparepartsDetailsCreateRequest() *AlibabaTmallSparepartsDetailsCreateAPIRequest {
	return &AlibabaTmallSparepartsDetailsCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallSparepartsDetailsCreateAPIRequest) Reset() {
	r._sparePartsDetailsSaveRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallSparepartsDetailsCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.tmall.spareparts.details.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallSparepartsDetailsCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallSparepartsDetailsCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSparePartsDetailsSaveRequest is SparePartsDetailsSaveRequest Setter
// 请求入参
func (r *AlibabaTmallSparepartsDetailsCreateAPIRequest) SetSparePartsDetailsSaveRequest(_sparePartsDetailsSaveRequest *SparePartsDetailsSaveRequest) error {
	r._sparePartsDetailsSaveRequest = _sparePartsDetailsSaveRequest
	r.Set("spare_parts_details_save_request", _sparePartsDetailsSaveRequest)
	return nil
}

// GetSparePartsDetailsSaveRequest SparePartsDetailsSaveRequest Getter
func (r AlibabaTmallSparepartsDetailsCreateAPIRequest) GetSparePartsDetailsSaveRequest() *SparePartsDetailsSaveRequest {
	return r._sparePartsDetailsSaveRequest
}

var poolAlibabaTmallSparepartsDetailsCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallSparepartsDetailsCreateRequest()
	},
}

// GetAlibabaTmallSparepartsDetailsCreateRequest 从 sync.Pool 获取 AlibabaTmallSparepartsDetailsCreateAPIRequest
func GetAlibabaTmallSparepartsDetailsCreateAPIRequest() *AlibabaTmallSparepartsDetailsCreateAPIRequest {
	return poolAlibabaTmallSparepartsDetailsCreateAPIRequest.Get().(*AlibabaTmallSparepartsDetailsCreateAPIRequest)
}

// ReleaseAlibabaTmallSparepartsDetailsCreateAPIRequest 将 AlibabaTmallSparepartsDetailsCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallSparepartsDetailsCreateAPIRequest(v *AlibabaTmallSparepartsDetailsCreateAPIRequest) {
	v.Reset()
	poolAlibabaTmallSparepartsDetailsCreateAPIRequest.Put(v)
}
