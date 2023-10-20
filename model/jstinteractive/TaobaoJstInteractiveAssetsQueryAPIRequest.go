package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstinteractiveassetsqueryAPIRequest 查询可配置任务素材接口 API请求
// taobao.jst.interactive.assets.query
//
// 查询可配置任务素材列表，用以配置任务素材
type TaobaojstinteractiveassetsqueryAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
}

// NewTaobaojstinteractiveassetsqueryRequest 初始化TaobaojstinteractiveassetsqueryAPIRequest对象
func NewTaobaojstinteractiveassetsqueryRequest() *TaobaojstinteractiveassetsqueryAPIRequest {
	return &TaobaojstinteractiveassetsqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstinteractiveassetsqueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.assets.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstinteractiveassetsqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstinteractiveassetsqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniAppId is MiniAppId Setter
// 小程序id
func (r *TaobaojstinteractiveassetsqueryAPIRequest) SetMiniAppId(_miniAppId string) error {
	r._miniAppId = _miniAppId
	r.Set("mini_app_id", _miniAppId)
	return nil
}

// GetMiniAppId MiniAppId Getter
func (r TaobaojstinteractiveassetsqueryAPIRequest) GetMiniAppId() string {
	return r._miniAppId
}
