package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlegoosefishuserinfoqueryAPIRequest 闲鱼三方容器用户信息获取 API请求
// alibaba.idle.goosefish.user.info.query
//
// 闲鱼三方容器用户信息获取
type AlibabaidlegoosefishuserinfoqueryAPIRequest struct {
	model.Params
}

// NewAlibabaidlegoosefishuserinfoqueryRequest 初始化AlibabaidlegoosefishuserinfoqueryAPIRequest对象
func NewAlibabaidlegoosefishuserinfoqueryRequest() *AlibabaidlegoosefishuserinfoqueryAPIRequest {
	return &AlibabaidlegoosefishuserinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlegoosefishuserinfoqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.goosefish.user.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlegoosefishuserinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlegoosefishuserinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
