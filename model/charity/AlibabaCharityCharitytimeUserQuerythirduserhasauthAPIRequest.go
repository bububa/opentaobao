package charity

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest 查询是否绑定3小时账号 API请求
// alibaba.charity.charitytime.user.querythirduserhasauth
//
// 查询是否绑定3小时账号
type AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest struct {
	model.Params
	// 参数对象
	_queryThirdUserHasAuthHsfRequest *QueryThirdUserHasAuthHsfRequest
}

// NewAlibabaCharityCharitytimeUserQuerythirduserhasauthRequest 初始化AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest对象
func NewAlibabaCharityCharitytimeUserQuerythirduserhasauthRequest() *AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest {
	return &AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest) Reset() {
	r._queryThirdUserHasAuthHsfRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.charitytime.user.querythirduserhasauth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryThirdUserHasAuthHsfRequest is QueryThirdUserHasAuthHsfRequest Setter
// 参数对象
func (r *AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest) SetQueryThirdUserHasAuthHsfRequest(_queryThirdUserHasAuthHsfRequest *QueryThirdUserHasAuthHsfRequest) error {
	r._queryThirdUserHasAuthHsfRequest = _queryThirdUserHasAuthHsfRequest
	r.Set("query_third_user_has_auth_hsf_request", _queryThirdUserHasAuthHsfRequest)
	return nil
}

// GetQueryThirdUserHasAuthHsfRequest QueryThirdUserHasAuthHsfRequest Getter
func (r AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest) GetQueryThirdUserHasAuthHsfRequest() *QueryThirdUserHasAuthHsfRequest {
	return r._queryThirdUserHasAuthHsfRequest
}

var poolAlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCharityCharitytimeUserQuerythirduserhasauthRequest()
	},
}

// GetAlibabaCharityCharitytimeUserQuerythirduserhasauthRequest 从 sync.Pool 获取 AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest
func GetAlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest() *AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest {
	return poolAlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest.Get().(*AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest)
}

// ReleaseAlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest 将 AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest 放入 sync.Pool
func ReleaseAlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest(v *AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest) {
	v.Reset()
	poolAlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest.Put(v)
}
