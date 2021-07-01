package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIRequest
【已废除】同步历史的销售数据 API请求
alibaba.tmallgenie.scp.plan.history.sale.qty.get

同步历史的销售数据 */
type AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabaTmallgenieScpPlanHistorySaleQtyGetRequest 初始化AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIRequest对象
func NewAlibabaTmallgenieScpPlanHistorySaleQtyGetRequest() *AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIRequest {
	return &AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.history.sale.qty.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// Get RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}
