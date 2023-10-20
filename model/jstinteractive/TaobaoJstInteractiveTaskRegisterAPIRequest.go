package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstinteractivetaskregisterAPIRequest 互动任务开通接口 API请求
// taobao.jst.interactive.task.register
//
// 调用互动任务开通接口为小程序开通互动任务
type TaobaojstinteractivetaskregisterAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
}

// NewTaobaojstinteractivetaskregisterRequest 初始化TaobaojstinteractivetaskregisterAPIRequest对象
func NewTaobaojstinteractivetaskregisterRequest() *TaobaojstinteractivetaskregisterAPIRequest {
	return &TaobaojstinteractivetaskregisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstinteractivetaskregisterAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.task.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstinteractivetaskregisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstinteractivetaskregisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniAppId is MiniAppId Setter
// 小程序id
func (r *TaobaojstinteractivetaskregisterAPIRequest) SetMiniAppId(_miniAppId string) error {
	r._miniAppId = _miniAppId
	r.Set("mini_app_id", _miniAppId)
	return nil
}

// GetMiniAppId MiniAppId Getter
func (r TaobaojstinteractivetaskregisterAPIRequest) GetMiniAppId() string {
	return r._miniAppId
}
