package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplannetdemandrawuploadAPIRequest 二级物料净需求回传(TL+1) API请求
// alibaba.tmallgenie.scp.plan.netdemand.raw.upload
//
// 二级物料净需求回传(TL+1)
type AlibabatmallgeniescpplannetdemandrawuploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRawRequest *NetDemandRawRequest
}

// NewAlibabatmallgeniescpplannetdemandrawuploadRequest 初始化AlibabatmallgeniescpplannetdemandrawuploadAPIRequest对象
func NewAlibabatmallgeniescpplannetdemandrawuploadRequest() *AlibabatmallgeniescpplannetdemandrawuploadAPIRequest {
	return &AlibabatmallgeniescpplannetdemandrawuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplannetdemandrawuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.netdemand.raw.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplannetdemandrawuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplannetdemandrawuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNetDemandRawRequest is NetDemandRawRequest Setter
// 对象
func (r *AlibabatmallgeniescpplannetdemandrawuploadAPIRequest) SetNetDemandRawRequest(_netDemandRawRequest *NetDemandRawRequest) error {
	r._netDemandRawRequest = _netDemandRawRequest
	r.Set("net_demand_raw_request", _netDemandRawRequest)
	return nil
}

// GetNetDemandRawRequest NetDemandRawRequest Getter
func (r AlibabatmallgeniescpplannetdemandrawuploadAPIRequest) GetNetDemandRawRequest() *NetDemandRawRequest {
	return r._netDemandRawRequest
}
