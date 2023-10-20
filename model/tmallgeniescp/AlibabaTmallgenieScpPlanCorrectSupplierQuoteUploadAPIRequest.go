package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplancorrectsupplierquoteuploadAPIRequest 供应商校准后的配额同步 API请求
// alibaba.tmallgenie.scp.plan.correct.supplier.quote.upload
//
// 供应商校准后的配额同步
type AlibabatmallgeniescpplancorrectsupplierquoteuploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// NewAlibabatmallgeniescpplancorrectsupplierquoteuploadRequest 初始化AlibabatmallgeniescpplancorrectsupplierquoteuploadAPIRequest对象
func NewAlibabatmallgeniescpplancorrectsupplierquoteuploadRequest() *AlibabatmallgeniescpplancorrectsupplierquoteuploadAPIRequest {
	return &AlibabatmallgeniescpplancorrectsupplierquoteuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplancorrectsupplierquoteuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.correct.supplier.quote.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplancorrectsupplierquoteuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplancorrectsupplierquoteuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNetDemandRequest is NetDemandRequest Setter
// 对象
func (r *AlibabatmallgeniescpplancorrectsupplierquoteuploadAPIRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
	r._netDemandRequest = _netDemandRequest
	r.Set("net_demand_request", _netDemandRequest)
	return nil
}

// GetNetDemandRequest NetDemandRequest Getter
func (r AlibabatmallgeniescpplancorrectsupplierquoteuploadAPIRequest) GetNetDemandRequest() *NetDemandRequest {
	return r._netDemandRequest
}
