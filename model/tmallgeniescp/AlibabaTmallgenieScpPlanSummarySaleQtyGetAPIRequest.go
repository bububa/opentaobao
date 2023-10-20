package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIRequest 同步销售数据按照渠道类型汇总 API请求
// alibaba.tmallgenie.scp.plan.summary.sale.qty.get
//
// 同步销售数据按照渠道类型汇总
type AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabaTmallgenieScpPlanSummarySaleQtyGetRequest 初始化AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIRequest对象
func NewAlibabaTmallgenieScpPlanSummarySaleQtyGetRequest() *AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIRequest {
	return &AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.summary.sale.qty.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}
