package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest 23-Net Demand净需求回传 API请求
// alibaba.tmallgenie.scp.plan.netdemand.upload
//
// Net Demand净需求回传
type AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// NewAlibabaTmallgenieScpPlanNetdemandUploadRequest 初始化AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanNetdemandUploadRequest() *AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.netdemand.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetNetDemandRequest is NetDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
	r._netDemandRequest = _netDemandRequest
	r.Set("net_demand_request", _netDemandRequest)
	return nil
}

// GetNetDemandRequest NetDemandRequest Getter
func (r AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest) GetNetDemandRequest() *NetDemandRequest {
	return r._netDemandRequest
}
