package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest 二级物料-LT内的POGAP数据回传 API请求
// alibaba.tmallgenie.scp.plan.rawpo.gap.return
//
// 二级物料-LT内的POGAP数据回传
type AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest struct {
	model.Params
	// 请求对象
	_rawPogapRequest *RawPurchaseOrderGapRequest
}

// NewAlibabaTmallgenieScpPlanRawpoGapReturnRequest 初始化AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest对象
func NewAlibabaTmallgenieScpPlanRawpoGapReturnRequest() *AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest {
	return &AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest) Reset() {
	r._rawPogapRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.rawpo.gap.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRawPogapRequest is RawPogapRequest Setter
// 请求对象
func (r *AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest) SetRawPogapRequest(_rawPogapRequest *RawPurchaseOrderGapRequest) error {
	r._rawPogapRequest = _rawPogapRequest
	r.Set("raw_pogap_request", _rawPogapRequest)
	return nil
}

// GetRawPogapRequest RawPogapRequest Getter
func (r AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest) GetRawPogapRequest() *RawPurchaseOrderGapRequest {
	return r._rawPogapRequest
}

var poolAlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanRawpoGapReturnRequest()
	},
}

// GetAlibabaTmallgenieScpPlanRawpoGapReturnRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest
func GetAlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest() *AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest {
	return poolAlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest.Get().(*AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest 将 AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest(v *AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest.Put(v)
}
