package jae

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoaplatformweakgetAPIRequest 活动平台弱登录接口 API请求
// taobao.aplatform.weakget
//
// 无线活动平台的开放接口，提供商品信息等的读操作
type TaobaoaplatformweakgetAPIRequest struct {
	model.Params
	// 客户端自带参数
	_paramRichClientInfo *RichClientInfo
	// 业务自定义参数
	_paramDto *ParamDto
}

// NewTaobaoaplatformweakgetRequest 初始化TaobaoaplatformweakgetAPIRequest对象
func NewTaobaoaplatformweakgetRequest() *TaobaoaplatformweakgetAPIRequest {
	return &TaobaoaplatformweakgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoaplatformweakgetAPIRequest) GetApiMethodName() string {
	return "taobao.aplatform.weakget"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoaplatformweakgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoaplatformweakgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRichClientInfo is ParamRichClientInfo Setter
// 客户端自带参数
func (r *TaobaoaplatformweakgetAPIRequest) SetParamRichClientInfo(_paramRichClientInfo *RichClientInfo) error {
	r._paramRichClientInfo = _paramRichClientInfo
	r.Set("param_rich_client_info", _paramRichClientInfo)
	return nil
}

// GetParamRichClientInfo ParamRichClientInfo Getter
func (r TaobaoaplatformweakgetAPIRequest) GetParamRichClientInfo() *RichClientInfo {
	return r._paramRichClientInfo
}

// SetParamDto is ParamDto Setter
// 业务自定义参数
func (r *TaobaoaplatformweakgetAPIRequest) SetParamDto(_paramDto *ParamDto) error {
	r._paramDto = _paramDto
	r.Set("param_dto", _paramDto)
	return nil
}

// GetParamDto ParamDto Getter
func (r TaobaoaplatformweakgetAPIRequest) GetParamDto() *ParamDto {
	return r._paramDto
}
