package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanchannelquoteuploadAPIRequest 9.1-同步渠道配额 API请求
// alibaba.tmallgenie.scp.plan.channel.quote.upload
//
// 同步渠道配额
type AlibabatmallgeniescpplanchannelquoteuploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// NewAlibabatmallgeniescpplanchannelquoteuploadRequest 初始化AlibabatmallgeniescpplanchannelquoteuploadAPIRequest对象
func NewAlibabatmallgeniescpplanchannelquoteuploadRequest() *AlibabatmallgeniescpplanchannelquoteuploadAPIRequest {
	return &AlibabatmallgeniescpplanchannelquoteuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplanchannelquoteuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.channel.quote.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplanchannelquoteuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplanchannelquoteuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNetDemandRequest is NetDemandRequest Setter
// 对象
func (r *AlibabatmallgeniescpplanchannelquoteuploadAPIRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
	r._netDemandRequest = _netDemandRequest
	r.Set("net_demand_request", _netDemandRequest)
	return nil
}

// GetNetDemandRequest NetDemandRequest Getter
func (r AlibabatmallgeniescpplanchannelquoteuploadAPIRequest) GetNetDemandRequest() *NetDemandRequest {
	return r._netDemandRequest
}
