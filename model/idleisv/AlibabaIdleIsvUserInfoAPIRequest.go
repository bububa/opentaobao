package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvuserinfoAPIRequest 闲鱼用户信息查询接口 API请求
// alibaba.idle.isv.user.info
//
// 闲鱼用户信息查询接口
type AlibabaidleisvuserinfoAPIRequest struct {
	model.Params
}

// NewAlibabaidleisvuserinfoRequest 初始化AlibabaidleisvuserinfoAPIRequest对象
func NewAlibabaidleisvuserinfoRequest() *AlibabaidleisvuserinfoAPIRequest {
	return &AlibabaidleisvuserinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleisvuserinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.user.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleisvuserinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleisvuserinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}
