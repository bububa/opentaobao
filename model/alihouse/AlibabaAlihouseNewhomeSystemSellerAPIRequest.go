package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomesystemsellerAPIRequest 商品发布授权 API请求
// alibaba.alihouse.newhome.system.seller
//
// 商品发布授权
type AlibabaalihousenewhomesystemsellerAPIRequest struct {
	model.Params
	// 请求实体
	_tokenCreateDto *TokenCreateDto
}

// NewAlibabaalihousenewhomesystemsellerRequest 初始化AlibabaalihousenewhomesystemsellerAPIRequest对象
func NewAlibabaalihousenewhomesystemsellerRequest() *AlibabaalihousenewhomesystemsellerAPIRequest {
	return &AlibabaalihousenewhomesystemsellerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomesystemsellerAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.system.seller"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomesystemsellerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomesystemsellerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTokenCreateDto is TokenCreateDto Setter
// 请求实体
func (r *AlibabaalihousenewhomesystemsellerAPIRequest) SetTokenCreateDto(_tokenCreateDto *TokenCreateDto) error {
	r._tokenCreateDto = _tokenCreateDto
	r.Set("token_create_dto", _tokenCreateDto)
	return nil
}

// GetTokenCreateDto TokenCreateDto Getter
func (r AlibabaalihousenewhomesystemsellerAPIRequest) GetTokenCreateDto() *TokenCreateDto {
	return r._tokenCreateDto
}
