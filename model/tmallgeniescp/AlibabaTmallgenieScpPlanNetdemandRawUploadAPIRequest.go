package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest 二级物料净需求回传(TL+1) API请求
// alibaba.tmallgenie.scp.plan.netdemand.raw.upload
//
// 二级物料净需求回传(TL+1)
type AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRawRequest *NetDemandRawRequest
}

// NewAlibabaTmallgenieScpPlanNetdemandRawUploadRequest 初始化AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanNetdemandRawUploadRequest() *AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.netdemand.raw.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetNetDemandRawRequest is NetDemandRawRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest) SetNetDemandRawRequest(_netDemandRawRequest *NetDemandRawRequest) error {
	r._netDemandRawRequest = _netDemandRawRequest
	r.Set("net_demand_raw_request", _netDemandRawRequest)
	return nil
}

// GetNetDemandRawRequest NetDemandRawRequest Getter
func (r AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest) GetNetDemandRawRequest() *NetDemandRawRequest {
	return r._netDemandRawRequest
}
