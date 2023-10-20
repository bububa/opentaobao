package ju

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJhsCommunityWechatLoginAPIRequest 聚划算用增淘外社群登录 API请求
// alibaba.jhs.community.wechat.login
//
// 聚划算用增淘外社群登录
type AlibabaJhsCommunityWechatLoginAPIRequest struct {
	model.Params
	// 动态令牌
	_code string
}

// NewAlibabaJhsCommunityWechatLoginRequest 初始化AlibabaJhsCommunityWechatLoginAPIRequest对象
func NewAlibabaJhsCommunityWechatLoginRequest() *AlibabaJhsCommunityWechatLoginAPIRequest {
	return &AlibabaJhsCommunityWechatLoginAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJhsCommunityWechatLoginAPIRequest) Reset() {
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJhsCommunityWechatLoginAPIRequest) GetApiMethodName() string {
	return "alibaba.jhs.community.wechat.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJhsCommunityWechatLoginAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJhsCommunityWechatLoginAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 动态令牌
func (r *AlibabaJhsCommunityWechatLoginAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaJhsCommunityWechatLoginAPIRequest) GetCode() string {
	return r._code
}

var poolAlibabaJhsCommunityWechatLoginAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJhsCommunityWechatLoginRequest()
	},
}

// GetAlibabaJhsCommunityWechatLoginRequest 从 sync.Pool 获取 AlibabaJhsCommunityWechatLoginAPIRequest
func GetAlibabaJhsCommunityWechatLoginAPIRequest() *AlibabaJhsCommunityWechatLoginAPIRequest {
	return poolAlibabaJhsCommunityWechatLoginAPIRequest.Get().(*AlibabaJhsCommunityWechatLoginAPIRequest)
}

// ReleaseAlibabaJhsCommunityWechatLoginAPIRequest 将 AlibabaJhsCommunityWechatLoginAPIRequest 放入 sync.Pool
func ReleaseAlibabaJhsCommunityWechatLoginAPIRequest(v *AlibabaJhsCommunityWechatLoginAPIRequest) {
	v.Reset()
	poolAlibabaJhsCommunityWechatLoginAPIRequest.Put(v)
}
