package tmallgeniescp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest) Reset() {
	r._requestExtendJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.current.po.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}

var poolAlibabaTmallgenieScpPlanCurrentPoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanCurrentPoGetRequest()
	},
}

// GetAlibabaTmallgenieScpPlanCurrentPoGetRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest
func GetAlibabaTmallgenieScpPlanCurrentPoGetAPIRequest() *AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest {
	return poolAlibabaTmallgenieScpPlanCurrentPoGetAPIRequest.Get().(*AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanCurrentPoGetAPIRequest 将 AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanCurrentPoGetAPIRequest(v *AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanCurrentPoGetAPIRequest.Put(v)
}
