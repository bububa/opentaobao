package jstinteractive

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractiveAssetsConfigAPIRequest 任务素材配置接口 API请求
// taobao.jst.interactive.assets.config
//
// 任务素材配置接口
type TaobaoJstInteractiveAssetsConfigAPIRequest struct {
	model.Params
	// []
	_assetsConfigList []AssetsConfig
	// 小程序id
	_miniAppId string
}

// NewTaobaoJstInteractiveAssetsConfigRequest 初始化TaobaoJstInteractiveAssetsConfigAPIRequest对象
func NewTaobaoJstInteractiveAssetsConfigRequest() *TaobaoJstInteractiveAssetsConfigAPIRequest {
	return &TaobaoJstInteractiveAssetsConfigAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstInteractiveAssetsConfigAPIRequest) Reset() {
	r._assetsConfigList = r._assetsConfigList[:0]
	r._miniAppId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveAssetsConfigAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.assets.config"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveAssetsConfigAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstInteractiveAssetsConfigAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAssetsConfigList is AssetsConfigList Setter
// []
func (r *TaobaoJstInteractiveAssetsConfigAPIRequest) SetAssetsConfigList(_assetsConfigList []AssetsConfig) error {
	r._assetsConfigList = _assetsConfigList
	r.Set("assets_config_list", _assetsConfigList)
	return nil
}

// GetAssetsConfigList AssetsConfigList Getter
func (r TaobaoJstInteractiveAssetsConfigAPIRequest) GetAssetsConfigList() []AssetsConfig {
	return r._assetsConfigList
}

// SetMiniAppId is MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveAssetsConfigAPIRequest) SetMiniAppId(_miniAppId string) error {
	r._miniAppId = _miniAppId
	r.Set("mini_app_id", _miniAppId)
	return nil
}

// GetMiniAppId MiniAppId Getter
func (r TaobaoJstInteractiveAssetsConfigAPIRequest) GetMiniAppId() string {
	return r._miniAppId
}

var poolTaobaoJstInteractiveAssetsConfigAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstInteractiveAssetsConfigRequest()
	},
}

// GetTaobaoJstInteractiveAssetsConfigRequest 从 sync.Pool 获取 TaobaoJstInteractiveAssetsConfigAPIRequest
func GetTaobaoJstInteractiveAssetsConfigAPIRequest() *TaobaoJstInteractiveAssetsConfigAPIRequest {
	return poolTaobaoJstInteractiveAssetsConfigAPIRequest.Get().(*TaobaoJstInteractiveAssetsConfigAPIRequest)
}

// ReleaseTaobaoJstInteractiveAssetsConfigAPIRequest 将 TaobaoJstInteractiveAssetsConfigAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstInteractiveAssetsConfigAPIRequest(v *TaobaoJstInteractiveAssetsConfigAPIRequest) {
	v.Reset()
	poolTaobaoJstInteractiveAssetsConfigAPIRequest.Put(v)
}
