package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrpsubmitAPIRequest 聚安全实人认证提交认证接口 API请求
// alibaba.security.jaq.rp.submit
//
// 聚安全实人认证提交认证接口
type AlibabasecurityjaqrpsubmitAPIRequest struct {
	model.Params
	// 认证token
	_verifyToken string
}

// NewAlibabasecurityjaqrpsubmitRequest 初始化AlibabasecurityjaqrpsubmitAPIRequest对象
func NewAlibabasecurityjaqrpsubmitRequest() *AlibabasecurityjaqrpsubmitAPIRequest {
	return &AlibabasecurityjaqrpsubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqrpsubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqrpsubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqrpsubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// 认证token
func (r *AlibabasecurityjaqrpsubmitAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabasecurityjaqrpsubmitAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}
