package tmallgeniescp

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.location.quote.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is NetDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
	r._netDemandRequest = _netDemandRequest
	r.Set("net_demand_request", _netDemandRequest)
	return nil
}

// Get NetDemandRequest Getter
func (r AlibabaTmallgenieScpPlanLocationQuoteUploadAPIRequest) GetNetDemandRequest() *NetDemandRequest {
	return r._netDemandRequest
}
