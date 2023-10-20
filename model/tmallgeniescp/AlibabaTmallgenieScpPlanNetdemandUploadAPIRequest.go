package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplannetdemanduploadAPIRequest 23-Net Demand净需求回传 API请求
// alibaba.tmallgenie.scp.plan.netdemand.upload
//
// Net Demand净需求回传
type AlibabatmallgeniescpplannetdemanduploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// NewAlibabatmallgeniescpplannetdemanduploadRequest 初始化AlibabatmallgeniescpplannetdemanduploadAPIRequest对象
func NewAlibabatmallgeniescpplannetdemanduploadRequest() *AlibabatmallgeniescpplannetdemanduploadAPIRequest {
	return &AlibabatmallgeniescpplannetdemanduploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplannetdemanduploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.netdemand.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplannetdemanduploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplannetdemanduploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNetDemandRequest is NetDemandRequest Setter
// 对象
func (r *AlibabatmallgeniescpplannetdemanduploadAPIRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
	r._netDemandRequest = _netDemandRequest
	r.Set("net_demand_request", _netDemandRequest)
	return nil
}

// GetNetDemandRequest NetDemandRequest Getter
func (r AlibabatmallgeniescpplannetdemanduploadAPIRequest) GetNetDemandRequest() *NetDemandRequest {
	return r._netDemandRequest
}
