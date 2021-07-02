package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest 二级物料-LT内的POGAP数据回传 API请求
// alibaba.tmallgenie.scp.plan.rawpo.gap.return
//
// 二级物料-LT内的POGAP数据回传
type AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest struct {
	model.Params
	// 请求对象
	_rawPogapRequest *RawPurchaseOrderGapRequest
}

// NewAlibabaTmallgenieScpPlanRawpoGapReturnRequest 初始化AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest对象
func NewAlibabaTmallgenieScpPlanRawpoGapReturnRequest() *AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest {
	return &AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.rawpo.gap.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRawPogapRequest is RawPogapRequest Setter
// 请求对象
func (r *AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest) SetRawPogapRequest(_rawPogapRequest *RawPurchaseOrderGapRequest) error {
	r._rawPogapRequest = _rawPogapRequest
	r.Set("raw_pogap_request", _rawPogapRequest)
	return nil
}

// GetRawPogapRequest RawPogapRequest Getter
func (r AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest) GetRawPogapRequest() *RawPurchaseOrderGapRequest {
	return r._rawPogapRequest
}
