package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractiveTaskRegisterAPIRequest 互动任务开通接口 API请求
// taobao.jst.interactive.task.register
//
// 调用互动任务开通接口为小程序开通互动任务
type TaobaoJstInteractiveTaskRegisterAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
}

// NewTaobaoJstInteractiveTaskRegisterRequest 初始化TaobaoJstInteractiveTaskRegisterAPIRequest对象
func NewTaobaoJstInteractiveTaskRegisterRequest() *TaobaoJstInteractiveTaskRegisterAPIRequest {
	return &TaobaoJstInteractiveTaskRegisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveTaskRegisterAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.task.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveTaskRegisterAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMiniAppId is MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveTaskRegisterAPIRequest) SetMiniAppId(_miniAppId string) error {
	r._miniAppId = _miniAppId
	r.Set("mini_app_id", _miniAppId)
	return nil
}

// GetMiniAppId MiniAppId Getter
func (r TaobaoJstInteractiveTaskRegisterAPIRequest) GetMiniAppId() string {
	return r._miniAppId
}
