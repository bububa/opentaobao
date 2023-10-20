package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanskusupplierquoteuploadAPIRequest 标准供应商配额同步 API请求
// alibaba.tmallgenie.scp.plan.sku.supplier.quote.upload
//
// 标准供应商配额同步
type AlibabatmallgeniescpplanskusupplierquoteuploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// NewAlibabatmallgeniescpplanskusupplierquoteuploadRequest 初始化AlibabatmallgeniescpplanskusupplierquoteuploadAPIRequest对象
func NewAlibabatmallgeniescpplanskusupplierquoteuploadRequest() *AlibabatmallgeniescpplanskusupplierquoteuploadAPIRequest {
	return &AlibabatmallgeniescpplanskusupplierquoteuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplanskusupplierquoteuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.sku.supplier.quote.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplanskusupplierquoteuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplanskusupplierquoteuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNetDemandRequest is NetDemandRequest Setter
// 对象
func (r *AlibabatmallgeniescpplanskusupplierquoteuploadAPIRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
	r._netDemandRequest = _netDemandRequest
	r.Set("net_demand_request", _netDemandRequest)
	return nil
}

// GetNetDemandRequest NetDemandRequest Getter
func (r AlibabatmallgeniescpplanskusupplierquoteuploadAPIRequest) GetNetDemandRequest() *NetDemandRequest {
	return r._netDemandRequest
}
