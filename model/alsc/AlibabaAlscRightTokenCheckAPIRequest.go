package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscRightTokenCheckAPIRequest 实物奖品凭证校验 API请求
// alibaba.alsc.right.token.check
//
// 实物奖品凭证校验
type AlibabaAlscRightTokenCheckAPIRequest struct {
	model.Params
	// 实物奖品凭证
	_token string
}

// NewAlibabaAlscRightTokenCheckRequest 初始化AlibabaAlscRightTokenCheckAPIRequest对象
func NewAlibabaAlscRightTokenCheckRequest() *AlibabaAlscRightTokenCheckAPIRequest {
	return &AlibabaAlscRightTokenCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscRightTokenCheckAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.right.token.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscRightTokenCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscRightTokenCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 实物奖品凭证
func (r *AlibabaAlscRightTokenCheckAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaAlscRightTokenCheckAPIRequest) GetToken() string {
	return r._token
}
