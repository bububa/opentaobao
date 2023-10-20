package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaguardaccessauthAPIRequest 鉴权 API请求
// alibaba.guard.access.auth
//
// 刷卡鉴权
type AlibabaguardaccessauthAPIRequest struct {
	model.Params
	// 鉴权结果DTO
	_identifyAuthDto *IdentifyAuthDto
}

// NewAlibabaguardaccessauthRequest 初始化AlibabaguardaccessauthAPIRequest对象
func NewAlibabaguardaccessauthRequest() *AlibabaguardaccessauthAPIRequest {
	return &AlibabaguardaccessauthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaguardaccessauthAPIRequest) GetApiMethodName() string {
	return "alibaba.guard.access.auth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaguardaccessauthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaguardaccessauthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdentifyAuthDto is IdentifyAuthDto Setter
// 鉴权结果DTO
func (r *AlibabaguardaccessauthAPIRequest) SetIdentifyAuthDto(_identifyAuthDto *IdentifyAuthDto) error {
	r._identifyAuthDto = _identifyAuthDto
	r.Set("identify_auth_dto", _identifyAuthDto)
	return nil
}

// GetIdentifyAuthDto IdentifyAuthDto Getter
func (r AlibabaguardaccessauthAPIRequest) GetIdentifyAuthDto() *IdentifyAuthDto {
	return r._identifyAuthDto
}
