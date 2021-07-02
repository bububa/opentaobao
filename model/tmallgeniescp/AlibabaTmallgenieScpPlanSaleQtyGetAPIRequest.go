package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest 12-同步销售数据 API请求
// alibaba.tmallgenie.scp.plan.sale.qty.get
//
// 同步销售数据
type AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabaTmallgenieScpPlanSaleQtyGetRequest 初始化AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest对象
func NewAlibabaTmallgenieScpPlanSaleQtyGetRequest() *AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest {
	return &AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.sale.qty.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}
