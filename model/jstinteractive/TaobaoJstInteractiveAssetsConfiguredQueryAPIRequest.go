package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest 查询已配置的任务素材列表接口 API请求
// taobao.jst.interactive.assets.configured.query
//
// 查询已配置任务素材列表
type TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
}

// NewTaobaoJstInteractiveAssetsConfiguredQueryRequest 初始化TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest对象
func NewTaobaoJstInteractiveAssetsConfiguredQueryRequest() *TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest {
	return &TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.assets.configured.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMiniAppId is MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest) SetMiniAppId(_miniAppId string) error {
	r._miniAppId = _miniAppId
	r.Set("mini_app_id", _miniAppId)
	return nil
}

// GetMiniAppId MiniAppId Getter
func (r TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest) GetMiniAppId() string {
	return r._miniAppId
}
