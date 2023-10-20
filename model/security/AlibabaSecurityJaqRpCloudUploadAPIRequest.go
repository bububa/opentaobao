package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrpclouduploadAPIRequest 实人认证云上传接口 API请求
// alibaba.security.jaq.rp.cloud.upload
//
// 聚安全实人认证上传认证信息
type AlibabasecurityjaqrpclouduploadAPIRequest struct {
	model.Params
	// []
	_elements []Elements
	// 认证token
	_verifyToken string
}

// NewAlibabasecurityjaqrpclouduploadRequest 初始化AlibabasecurityjaqrpclouduploadAPIRequest对象
func NewAlibabasecurityjaqrpclouduploadRequest() *AlibabasecurityjaqrpclouduploadAPIRequest {
	return &AlibabasecurityjaqrpclouduploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqrpclouduploadAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.cloud.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqrpclouduploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqrpclouduploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetElements is Elements Setter
// []
func (r *AlibabasecurityjaqrpclouduploadAPIRequest) SetElements(_elements []Elements) error {
	r._elements = _elements
	r.Set("elements", _elements)
	return nil
}

// GetElements Elements Getter
func (r AlibabasecurityjaqrpclouduploadAPIRequest) GetElements() []Elements {
	return r._elements
}

// SetVerifyToken is VerifyToken Setter
// 认证token
func (r *AlibabasecurityjaqrpclouduploadAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabasecurityjaqrpclouduploadAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}
