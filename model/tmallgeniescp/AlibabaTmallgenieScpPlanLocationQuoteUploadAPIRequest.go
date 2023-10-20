package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanlocationquoteuploadAPIRequest 9.2-同步地点配额 API请求
// alibaba.tmallgenie.scp.plan.location.quote.upload
//
// 同步地点配额
type AlibabatmallgeniescpplanlocationquoteuploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// NewAlibabatmallgeniescpplanlocationquoteuploadRequest 初始化AlibabatmallgeniescpplanlocationquoteuploadAPIRequest对象
func NewAlibabatmallgeniescpplanlocationquoteuploadRequest() *AlibabatmallgeniescpplanlocationquoteuploadAPIRequest {
	return &AlibabatmallgeniescpplanlocationquoteuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplanlocationquoteuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.location.quote.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplanlocationquoteuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplanlocationquoteuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNetDemandRequest is NetDemandRequest Setter
// 对象
func (r *AlibabatmallgeniescpplanlocationquoteuploadAPIRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
	r._netDemandRequest = _netDemandRequest
	r.Set("net_demand_request", _netDemandRequest)
	return nil
}

// GetNetDemandRequest NetDemandRequest Getter
func (r AlibabatmallgeniescpplanlocationquoteuploadAPIRequest) GetNetDemandRequest() *NetDemandRequest {
	return r._netDemandRequest
}
