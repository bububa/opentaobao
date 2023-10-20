package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstinteractiveassetsconfiguredqueryAPIRequest 查询已配置的任务素材列表接口 API请求
// taobao.jst.interactive.assets.configured.query
//
// 查询已配置任务素材列表
type TaobaojstinteractiveassetsconfiguredqueryAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
}

// NewTaobaojstinteractiveassetsconfiguredqueryRequest 初始化TaobaojstinteractiveassetsconfiguredqueryAPIRequest对象
func NewTaobaojstinteractiveassetsconfiguredqueryRequest() *TaobaojstinteractiveassetsconfiguredqueryAPIRequest {
	return &TaobaojstinteractiveassetsconfiguredqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstinteractiveassetsconfiguredqueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.assets.configured.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstinteractiveassetsconfiguredqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstinteractiveassetsconfiguredqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniAppId is MiniAppId Setter
// 小程序id
func (r *TaobaojstinteractiveassetsconfiguredqueryAPIRequest) SetMiniAppId(_miniAppId string) error {
	r._miniAppId = _miniAppId
	r.Set("mini_app_id", _miniAppId)
	return nil
}

// GetMiniAppId MiniAppId Getter
func (r TaobaojstinteractiveassetsconfiguredqueryAPIRequest) GetMiniAppId() string {
	return r._miniAppId
}
