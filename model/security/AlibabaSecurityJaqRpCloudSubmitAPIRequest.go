package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrpcloudsubmitAPIRequest 实人认证云服务提交接口 API请求
// alibaba.security.jaq.rp.cloud.submit
//
// 聚安全实人认证提交认证接口
type AlibabasecurityjaqrpcloudsubmitAPIRequest struct {
	model.Params
	// 认证token
	_verifyToken string
}

// NewAlibabasecurityjaqrpcloudsubmitRequest 初始化AlibabasecurityjaqrpcloudsubmitAPIRequest对象
func NewAlibabasecurityjaqrpcloudsubmitRequest() *AlibabasecurityjaqrpcloudsubmitAPIRequest {
	return &AlibabasecurityjaqrpcloudsubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqrpcloudsubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.cloud.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqrpcloudsubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqrpcloudsubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// 认证token
func (r *AlibabasecurityjaqrpcloudsubmitAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabasecurityjaqrpcloudsubmitAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}
