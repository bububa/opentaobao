package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeSystemSellerAPIRequest 商品发布授权 API请求
// alibaba.alihouse.newhome.system.seller
//
// 商品发布授权
type AlibabaAlihouseNewhomeSystemSellerAPIRequest struct {
	model.Params
	// 请求实体
	_tokenCreateDto *TokenCreateDto
}

// NewAlibabaAlihouseNewhomeSystemSellerRequest 初始化AlibabaAlihouseNewhomeSystemSellerAPIRequest对象
func NewAlibabaAlihouseNewhomeSystemSellerRequest() *AlibabaAlihouseNewhomeSystemSellerAPIRequest {
	return &AlibabaAlihouseNewhomeSystemSellerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeSystemSellerAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.system.seller"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeSystemSellerAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTokenCreateDto is TokenCreateDto Setter
// 请求实体
func (r *AlibabaAlihouseNewhomeSystemSellerAPIRequest) SetTokenCreateDto(_tokenCreateDto *TokenCreateDto) error {
	r._tokenCreateDto = _tokenCreateDto
	r.Set("token_create_dto", _tokenCreateDto)
	return nil
}

// GetTokenCreateDto TokenCreateDto Getter
func (r AlibabaAlihouseNewhomeSystemSellerAPIRequest) GetTokenCreateDto() *TokenCreateDto {
	return r._tokenCreateDto
}
