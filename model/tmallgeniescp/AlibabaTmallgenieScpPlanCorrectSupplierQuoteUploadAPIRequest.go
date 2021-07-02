package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest 供应商校准后的配额同步 API请求
// alibaba.tmallgenie.scp.plan.correct.supplier.quote.upload
//
// 供应商校准后的配额同步
type AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// NewAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest 初始化AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest() *AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.correct.supplier.quote.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetNetDemandRequest is NetDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
	r._netDemandRequest = _netDemandRequest
	r.Set("net_demand_request", _netDemandRequest)
	return nil
}

// GetNetDemandRequest NetDemandRequest Getter
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest) GetNetDemandRequest() *NetDemandRequest {
	return r._netDemandRequest
}
