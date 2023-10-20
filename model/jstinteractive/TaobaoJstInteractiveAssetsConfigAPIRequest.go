package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstinteractiveassetsconfigAPIRequest 任务素材配置接口 API请求
// taobao.jst.interactive.assets.config
//
// 任务素材配置接口
type TaobaojstinteractiveassetsconfigAPIRequest struct {
	model.Params
	// []
	_assetsConfigList []AssetsConfig
	// 小程序id
	_miniAppId string
}

// NewTaobaojstinteractiveassetsconfigRequest 初始化TaobaojstinteractiveassetsconfigAPIRequest对象
func NewTaobaojstinteractiveassetsconfigRequest() *TaobaojstinteractiveassetsconfigAPIRequest {
	return &TaobaojstinteractiveassetsconfigAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstinteractiveassetsconfigAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.assets.config"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstinteractiveassetsconfigAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstinteractiveassetsconfigAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAssetsConfigList is AssetsConfigList Setter
// []
func (r *TaobaojstinteractiveassetsconfigAPIRequest) SetAssetsConfigList(_assetsConfigList []AssetsConfig) error {
	r._assetsConfigList = _assetsConfigList
	r.Set("assets_config_list", _assetsConfigList)
	return nil
}

// GetAssetsConfigList AssetsConfigList Getter
func (r TaobaojstinteractiveassetsconfigAPIRequest) GetAssetsConfigList() []AssetsConfig {
	return r._assetsConfigList
}

// SetMiniAppId is MiniAppId Setter
// 小程序id
func (r *TaobaojstinteractiveassetsconfigAPIRequest) SetMiniAppId(_miniAppId string) error {
	r._miniAppId = _miniAppId
	r.Set("mini_app_id", _miniAppId)
	return nil
}

// GetMiniAppId MiniAppId Getter
func (r TaobaojstinteractiveassetsconfigAPIRequest) GetMiniAppId() string {
	return r._miniAppId
}
