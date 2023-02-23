package damaiticklet

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMxOpengatewayScriptAPIRequest 第三方剧本数据推送 API请求
// alibaba.damai.mx.opengateway.script
//
// 第三方剧本数据推送
type AlibabaDamaiMxOpengatewayScriptAPIRequest struct {
	model.Params
	// 接口入参
	_scriptInfoOpenParam *ScriptInfoOpenParam
}

// NewAlibabaDamaiMxOpengatewayScriptRequest 初始化AlibabaDamaiMxOpengatewayScriptAPIRequest对象
func NewAlibabaDamaiMxOpengatewayScriptRequest() *AlibabaDamaiMxOpengatewayScriptAPIRequest {
	return &AlibabaDamaiMxOpengatewayScriptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMxOpengatewayScriptAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mx.opengateway.script"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMxOpengatewayScriptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMxOpengatewayScriptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScriptInfoOpenParam is ScriptInfoOpenParam Setter
// 接口入参
func (r *AlibabaDamaiMxOpengatewayScriptAPIRequest) SetScriptInfoOpenParam(_scriptInfoOpenParam *ScriptInfoOpenParam) error {
	r._scriptInfoOpenParam = _scriptInfoOpenParam
	r.Set("script_info_open_param", _scriptInfoOpenParam)
	return nil
}

// GetScriptInfoOpenParam ScriptInfoOpenParam Getter
func (r AlibabaDamaiMxOpengatewayScriptAPIRequest) GetScriptInfoOpenParam() *ScriptInfoOpenParam {
	return r._scriptInfoOpenParam
}
