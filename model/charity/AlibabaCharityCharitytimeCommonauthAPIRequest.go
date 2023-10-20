package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacharitycharitytimecommonauthAPIRequest 通用授权 API请求
// alibaba.charity.charitytime.commonauth
//
// 三小时和外部账号绑定通用top 返回跳转链接进行绑定
type AlibabacharitycharitytimecommonauthAPIRequest struct {
	model.Params
	// 请求对象
	_jumpAddressHsfRequest *JumpAddressHsfRequest
}

// NewAlibabacharitycharitytimecommonauthRequest 初始化AlibabacharitycharitytimecommonauthAPIRequest对象
func NewAlibabacharitycharitytimecommonauthRequest() *AlibabacharitycharitytimecommonauthAPIRequest {
	return &AlibabacharitycharitytimecommonauthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacharitycharitytimecommonauthAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.charitytime.commonauth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacharitycharitytimecommonauthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacharitycharitytimecommonauthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetJumpAddressHsfRequest is JumpAddressHsfRequest Setter
// 请求对象
func (r *AlibabacharitycharitytimecommonauthAPIRequest) SetJumpAddressHsfRequest(_jumpAddressHsfRequest *JumpAddressHsfRequest) error {
	r._jumpAddressHsfRequest = _jumpAddressHsfRequest
	r.Set("jump_address_hsf_request", _jumpAddressHsfRequest)
	return nil
}

// GetJumpAddressHsfRequest JumpAddressHsfRequest Getter
func (r AlibabacharitycharitytimecommonauthAPIRequest) GetJumpAddressHsfRequest() *JumpAddressHsfRequest {
	return r._jumpAddressHsfRequest
}
