package damaiticklet

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimxopengatewayscriptAPIRequest 第三方剧本数据推送 API请求
// alibaba.damai.mx.opengateway.script
//
// 第三方剧本数据推送
type AlibabadamaimxopengatewayscriptAPIRequest struct {
	model.Params
	// 接口入参
	_scriptInfoOpenParam *ScriptInfoOpenParam
}

// NewAlibabadamaimxopengatewayscriptRequest 初始化AlibabadamaimxopengatewayscriptAPIRequest对象
func NewAlibabadamaimxopengatewayscriptRequest() *AlibabadamaimxopengatewayscriptAPIRequest {
	return &AlibabadamaimxopengatewayscriptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimxopengatewayscriptAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mx.opengateway.script"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimxopengatewayscriptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimxopengatewayscriptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScriptInfoOpenParam is ScriptInfoOpenParam Setter
// 接口入参
func (r *AlibabadamaimxopengatewayscriptAPIRequest) SetScriptInfoOpenParam(_scriptInfoOpenParam *ScriptInfoOpenParam) error {
	r._scriptInfoOpenParam = _scriptInfoOpenParam
	r.Set("script_info_open_param", _scriptInfoOpenParam)
	return nil
}

// GetScriptInfoOpenParam ScriptInfoOpenParam Getter
func (r AlibabadamaimxopengatewayscriptAPIRequest) GetScriptInfoOpenParam() *ScriptInfoOpenParam {
	return r._scriptInfoOpenParam
}
