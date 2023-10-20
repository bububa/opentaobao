package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvuserqueryAPIRequest 服务商ISV闲鱼用户信息查询 API请求
// alibaba.idle.isv.user.query
//
// 服务商ISV闲鱼用户信息查询
type AlibabaidleisvuserqueryAPIRequest struct {
	model.Params
}

// NewAlibabaidleisvuserqueryRequest 初始化AlibabaidleisvuserqueryAPIRequest对象
func NewAlibabaidleisvuserqueryRequest() *AlibabaidleisvuserqueryAPIRequest {
	return &AlibabaidleisvuserqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleisvuserqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.user.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleisvuserqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleisvuserqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
