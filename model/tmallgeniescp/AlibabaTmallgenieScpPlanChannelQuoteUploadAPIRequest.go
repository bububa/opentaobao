package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest 9.1-同步渠道配额 API请求
// alibaba.tmallgenie.scp.plan.channel.quote.upload
//
// 同步渠道配额
type AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// NewAlibabaTmallgenieScpPlanChannelQuoteUploadRequest 初始化AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanChannelQuoteUploadRequest() *AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest) Reset() {
	r._netDemandRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.channel.quote.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNetDemandRequest is NetDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
	r._netDemandRequest = _netDemandRequest
	r.Set("net_demand_request", _netDemandRequest)
	return nil
}

// GetNetDemandRequest NetDemandRequest Getter
func (r AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest) GetNetDemandRequest() *NetDemandRequest {
	return r._netDemandRequest
}

var poolAlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanChannelQuoteUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanChannelQuoteUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest
func GetAlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest() *AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest 将 AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest(v *AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest.Put(v)
}
