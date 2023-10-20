package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplancurrentpogetAPIRequest 11-同步本周的po单（从W-1周到W+4周） API请求
// alibaba.tmallgenie.scp.plan.current.po.get
//
// 11-同步本周的po单（从W-1周到W+4周）
type AlibabatmallgeniescpplancurrentpogetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabatmallgeniescpplancurrentpogetRequest 初始化AlibabatmallgeniescpplancurrentpogetAPIRequest对象
func NewAlibabatmallgeniescpplancurrentpogetRequest() *AlibabatmallgeniescpplancurrentpogetAPIRequest {
	return &AlibabatmallgeniescpplancurrentpogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplancurrentpogetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.current.po.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplancurrentpogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplancurrentpogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展参数
func (r *AlibabatmallgeniescpplancurrentpogetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabatmallgeniescpplancurrentpogetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}
