package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest 标准供应商配额同步 API请求
// alibaba.tmallgenie.scp.plan.sku.supplier.quote.upload
//
// 标准供应商配额同步
type AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// NewAlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest 初始化AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest() *AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest) Reset() {
	r._netDemandRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.sku.supplier.quote.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNetDemandRequest is NetDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
	r._netDemandRequest = _netDemandRequest
	r.Set("net_demand_request", _netDemandRequest)
	return nil
}

// GetNetDemandRequest NetDemandRequest Getter
func (r AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest) GetNetDemandRequest() *NetDemandRequest {
	return r._netDemandRequest
}

var poolAlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest
func GetAlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest() *AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest 将 AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest(v *AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest.Put(v)
}
