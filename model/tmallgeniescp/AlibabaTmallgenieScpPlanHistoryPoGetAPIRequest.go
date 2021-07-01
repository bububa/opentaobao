package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanHistoryPoGetAPIRequest
【已废除】11-同步历史所有的po单 API请求
alibaba.tmallgenie.scp.plan.history.po.get

同步历史po单 */
type AlibabaTmallgenieScpPlanHistoryPoGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabaTmallgenieScpPlanHistoryPoGetRequest 初始化AlibabaTmallgenieScpPlanHistoryPoGetAPIRequest对象
func NewAlibabaTmallgenieScpPlanHistoryPoGetRequest() *AlibabaTmallgenieScpPlanHistoryPoGetAPIRequest {
	return &AlibabaTmallgenieScpPlanHistoryPoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanHistoryPoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.history.po.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanHistoryPoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanHistoryPoGetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// Get RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanHistoryPoGetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}
