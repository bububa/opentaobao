package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrpqueryAPIRequest 聚安全实人认证查询认证结果 API请求
// alibaba.security.jaq.rp.query
//
// 聚安全实人认证查询认证结果
type AlibabasecurityjaqrpqueryAPIRequest struct {
	model.Params
	// token
	_verifyToken string
}

// NewAlibabasecurityjaqrpqueryRequest 初始化AlibabasecurityjaqrpqueryAPIRequest对象
func NewAlibabasecurityjaqrpqueryRequest() *AlibabasecurityjaqrpqueryAPIRequest {
	return &AlibabasecurityjaqrpqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqrpqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqrpqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqrpqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// token
func (r *AlibabasecurityjaqrpqueryAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabasecurityjaqrpqueryAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}
