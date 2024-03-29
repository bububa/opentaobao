package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest 9.2-同步地点配额 API请求
// alibaba.tmallgenie.scp.plan.location.quote.upload
//
// 同步地点配额
type AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// NewAlibabaTmallgenieScpPlanLocationQuoteUploadRequest 初始化AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanLocationQuoteUploadRequest() *AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest) Reset() {
	r._netDemandRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.location.quote.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNetDemandRequest is NetDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
	r._netDemandRequest = _netDemandRequest
	r.Set("net_demand_request", _netDemandRequest)
	return nil
}

// GetNetDemandRequest NetDemandRequest Getter
func (r AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest) GetNetDemandRequest() *NetDemandRequest {
	return r._netDemandRequest
}

var poolAlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanLocationQuoteUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanLocationQuoteUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest
func GetAlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest() *AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest 将 AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest(v *AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest.Put(v)
}
