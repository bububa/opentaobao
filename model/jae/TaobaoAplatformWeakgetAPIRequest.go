package jae

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAplatformWeakgetAPIRequest 活动平台弱登录接口 API请求
// taobao.aplatform.weakget
//
// 无线活动平台的开放接口，提供商品信息等的读操作
type TaobaoAplatformWeakgetAPIRequest struct {
	model.Params
	// 客户端自带参数
	_paramRichClientInfo *RichClientInfo
	// 业务自定义参数
	_paramDto *ParamDto
}

// NewTaobaoAplatformWeakgetRequest 初始化TaobaoAplatformWeakgetAPIRequest对象
func NewTaobaoAplatformWeakgetRequest() *TaobaoAplatformWeakgetAPIRequest {
	return &TaobaoAplatformWeakgetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAplatformWeakgetAPIRequest) Reset() {
	r._paramRichClientInfo = nil
	r._paramDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAplatformWeakgetAPIRequest) GetApiMethodName() string {
	return "taobao.aplatform.weakget"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAplatformWeakgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAplatformWeakgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRichClientInfo is ParamRichClientInfo Setter
// 客户端自带参数
func (r *TaobaoAplatformWeakgetAPIRequest) SetParamRichClientInfo(_paramRichClientInfo *RichClientInfo) error {
	r._paramRichClientInfo = _paramRichClientInfo
	r.Set("param_rich_client_info", _paramRichClientInfo)
	return nil
}

// GetParamRichClientInfo ParamRichClientInfo Getter
func (r TaobaoAplatformWeakgetAPIRequest) GetParamRichClientInfo() *RichClientInfo {
	return r._paramRichClientInfo
}

// SetParamDto is ParamDto Setter
// 业务自定义参数
func (r *TaobaoAplatformWeakgetAPIRequest) SetParamDto(_paramDto *ParamDto) error {
	r._paramDto = _paramDto
	r.Set("param_dto", _paramDto)
	return nil
}

// GetParamDto ParamDto Getter
func (r TaobaoAplatformWeakgetAPIRequest) GetParamDto() *ParamDto {
	return r._paramDto
}

var poolTaobaoAplatformWeakgetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAplatformWeakgetRequest()
	},
}

// GetTaobaoAplatformWeakgetRequest 从 sync.Pool 获取 TaobaoAplatformWeakgetAPIRequest
func GetTaobaoAplatformWeakgetAPIRequest() *TaobaoAplatformWeakgetAPIRequest {
	return poolTaobaoAplatformWeakgetAPIRequest.Get().(*TaobaoAplatformWeakgetAPIRequest)
}

// ReleaseTaobaoAplatformWeakgetAPIRequest 将 TaobaoAplatformWeakgetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAplatformWeakgetAPIRequest(v *TaobaoAplatformWeakgetAPIRequest) {
	v.Reset()
	poolTaobaoAplatformWeakgetAPIRequest.Put(v)
}
