package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacharitycharitytimeuserquerythirduserhasauthAPIRequest 查询是否绑定3小时账号 API请求
// alibaba.charity.charitytime.user.querythirduserhasauth
//
// 查询是否绑定3小时账号
type AlibabacharitycharitytimeuserquerythirduserhasauthAPIRequest struct {
	model.Params
	// 参数对象
	_queryThirdUserHasAuthHsfRequest *QueryThirdUserHasAuthHsfRequest
}

// NewAlibabacharitycharitytimeuserquerythirduserhasauthRequest 初始化AlibabacharitycharitytimeuserquerythirduserhasauthAPIRequest对象
func NewAlibabacharitycharitytimeuserquerythirduserhasauthRequest() *AlibabacharitycharitytimeuserquerythirduserhasauthAPIRequest {
	return &AlibabacharitycharitytimeuserquerythirduserhasauthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacharitycharitytimeuserquerythirduserhasauthAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.charitytime.user.querythirduserhasauth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacharitycharitytimeuserquerythirduserhasauthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacharitycharitytimeuserquerythirduserhasauthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryThirdUserHasAuthHsfRequest is QueryThirdUserHasAuthHsfRequest Setter
// 参数对象
func (r *AlibabacharitycharitytimeuserquerythirduserhasauthAPIRequest) SetQueryThirdUserHasAuthHsfRequest(_queryThirdUserHasAuthHsfRequest *QueryThirdUserHasAuthHsfRequest) error {
	r._queryThirdUserHasAuthHsfRequest = _queryThirdUserHasAuthHsfRequest
	r.Set("query_third_user_has_auth_hsf_request", _queryThirdUserHasAuthHsfRequest)
	return nil
}

// GetQueryThirdUserHasAuthHsfRequest QueryThirdUserHasAuthHsfRequest Getter
func (r AlibabacharitycharitytimeuserquerythirduserhasauthAPIRequest) GetQueryThirdUserHasAuthHsfRequest() *QueryThirdUserHasAuthHsfRequest {
	return r._queryThirdUserHasAuthHsfRequest
}
