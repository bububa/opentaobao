package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleGoosefishUserInfoQueryAPIRequest 闲鱼三方容器用户信息获取 API请求
// alibaba.idle.goosefish.user.info.query
//
// 闲鱼三方容器用户信息获取
type AlibabaIdleGoosefishUserInfoQueryAPIRequest struct {
	model.Params
}

// NewAlibabaIdleGoosefishUserInfoQueryRequest 初始化AlibabaIdleGoosefishUserInfoQueryAPIRequest对象
func NewAlibabaIdleGoosefishUserInfoQueryRequest() *AlibabaIdleGoosefishUserInfoQueryAPIRequest {
	return &AlibabaIdleGoosefishUserInfoQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleGoosefishUserInfoQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.goosefish.user.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleGoosefishUserInfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleGoosefishUserInfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
