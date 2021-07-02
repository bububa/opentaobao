package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest 11-同步本周的po单（从W-1周到W+4周） API请求
// alibaba.tmallgenie.scp.plan.current.po.get
//
// 11-同步本周的po单（从W-1周到W+4周）
type AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabaTmallgenieScpPlanCurrentPoGetRequest 初始化AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest对象
func NewAlibabaTmallgenieScpPlanCurrentPoGetRequest() *AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest {
	return &AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.current.po.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// Get RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}
