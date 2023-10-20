package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeSystemSellerAPIRequest) Reset() {
	r._tokenCreateDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeSystemSellerAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.system.seller"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeSystemSellerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeSystemSellerAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihouseNewhomeSystemSellerAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeSystemSellerRequest()
	},
}

// GetAlibabaAlihouseNewhomeSystemSellerRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeSystemSellerAPIRequest
func GetAlibabaAlihouseNewhomeSystemSellerAPIRequest() *AlibabaAlihouseNewhomeSystemSellerAPIRequest {
	return poolAlibabaAlihouseNewhomeSystemSellerAPIRequest.Get().(*AlibabaAlihouseNewhomeSystemSellerAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeSystemSellerAPIRequest 将 AlibabaAlihouseNewhomeSystemSellerAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeSystemSellerAPIRequest(v *AlibabaAlihouseNewhomeSystemSellerAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeSystemSellerAPIRequest.Put(v)
}
