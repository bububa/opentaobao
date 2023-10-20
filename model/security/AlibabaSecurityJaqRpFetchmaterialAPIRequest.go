package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrpfetchmaterialAPIRequest 聚安全实人认证获取结果接口 API请求
// alibaba.security.jaq.rp.fetchmaterial
//
// 聚安全实人认证获取结果接口
type AlibabasecurityjaqrpfetchmaterialAPIRequest struct {
	model.Params
	// 消息服务推送的key
	_securityKey string
}

// NewAlibabasecurityjaqrpfetchmaterialRequest 初始化AlibabasecurityjaqrpfetchmaterialAPIRequest对象
func NewAlibabasecurityjaqrpfetchmaterialRequest() *AlibabasecurityjaqrpfetchmaterialAPIRequest {
	return &AlibabasecurityjaqrpfetchmaterialAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqrpfetchmaterialAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.fetchmaterial"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqrpfetchmaterialAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqrpfetchmaterialAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSecurityKey is SecurityKey Setter
// 消息服务推送的key
func (r *AlibabasecurityjaqrpfetchmaterialAPIRequest) SetSecurityKey(_securityKey string) error {
	r._securityKey = _securityKey
	r.Set("security_key", _securityKey)
	return nil
}

// GetSecurityKey SecurityKey Getter
func (r AlibabasecurityjaqrpfetchmaterialAPIRequest) GetSecurityKey() string {
	return r._securityKey
}
