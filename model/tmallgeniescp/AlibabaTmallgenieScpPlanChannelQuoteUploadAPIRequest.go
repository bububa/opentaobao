package tmallgeniescp

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.channel.quote.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is NetDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
	r._netDemandRequest = _netDemandRequest
	r.Set("net_demand_request", _netDemandRequest)
	return nil
}

// Get NetDemandRequest Getter
func (r AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest) GetNetDemandRequest() *NetDemandRequest {
	return r._netDemandRequest
}
