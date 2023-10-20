package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstinteractiveactivityqueryAPIRequest 互动任务活动查询接口 API请求
// taobao.jst.interactive.activity.query
//
// 互动任务活动查询接口
type TaobaojstinteractiveactivityqueryAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
}

// NewTaobaojstinteractiveactivityqueryRequest 初始化TaobaojstinteractiveactivityqueryAPIRequest对象
func NewTaobaojstinteractiveactivityqueryRequest() *TaobaojstinteractiveactivityqueryAPIRequest {
	return &TaobaojstinteractiveactivityqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstinteractiveactivityqueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstinteractiveactivityqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstinteractiveactivityqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniAppId is MiniAppId Setter
// 小程序id
func (r *TaobaojstinteractiveactivityqueryAPIRequest) SetMiniAppId(_miniAppId string) error {
	r._miniAppId = _miniAppId
	r.Set("mini_app_id", _miniAppId)
	return nil
}

// GetMiniAppId MiniAppId Getter
func (r TaobaojstinteractiveactivityqueryAPIRequest) GetMiniAppId() string {
	return r._miniAppId
}
