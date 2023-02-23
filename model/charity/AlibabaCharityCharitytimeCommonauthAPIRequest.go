package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityCharitytimeCommonauthAPIRequest 通用授权 API请求
// alibaba.charity.charitytime.commonauth
//
// 三小时和外部账号绑定通用top 返回跳转链接进行绑定
type AlibabaCharityCharitytimeCommonauthAPIRequest struct {
	model.Params
	// 请求对象
	_jumpAddressHsfRequest *JumpAddressHsfRequest
}

// NewAlibabaCharityCharitytimeCommonauthRequest 初始化AlibabaCharityCharitytimeCommonauthAPIRequest对象
func NewAlibabaCharityCharitytimeCommonauthRequest() *AlibabaCharityCharitytimeCommonauthAPIRequest {
	return &AlibabaCharityCharitytimeCommonauthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCharityCharitytimeCommonauthAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.charitytime.commonauth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCharityCharitytimeCommonauthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCharityCharitytimeCommonauthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetJumpAddressHsfRequest is JumpAddressHsfRequest Setter
// 请求对象
func (r *AlibabaCharityCharitytimeCommonauthAPIRequest) SetJumpAddressHsfRequest(_jumpAddressHsfRequest *JumpAddressHsfRequest) error {
	r._jumpAddressHsfRequest = _jumpAddressHsfRequest
	r.Set("jump_address_hsf_request", _jumpAddressHsfRequest)
	return nil
}

// GetJumpAddressHsfRequest JumpAddressHsfRequest Getter
func (r AlibabaCharityCharitytimeCommonauthAPIRequest) GetJumpAddressHsfRequest() *JumpAddressHsfRequest {
	return r._jumpAddressHsfRequest
}
