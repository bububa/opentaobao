package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
任务素材配置接口 API请求
taobao.jst.interactive.assets.config

任务素材配置接口
*/
type TaobaoJstInteractiveAssetsConfigAPIRequest struct {
    model.Params
    // []
    _assetsConfigList   []AssetsConfig
    // 小程序id
    _miniAppId   string
}

// 初始化TaobaoJstInteractiveAssetsConfigAPIRequest对象
func NewTaobaoJstInteractiveAssetsConfigRequest() *TaobaoJstInteractiveAssetsConfigAPIRequest{
    return &TaobaoJstInteractiveAssetsConfigAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveAssetsConfigAPIRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.assets.config"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveAssetsConfigAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AssetsConfigList Setter
// []
func (r *TaobaoJstInteractiveAssetsConfigAPIRequest) SetAssetsConfigList(_assetsConfigList []AssetsConfig) error {
    r._assetsConfigList = _assetsConfigList
    r.Set("assets_config_list", _assetsConfigList)
    return nil
}

// AssetsConfigList Getter
func (r TaobaoJstInteractiveAssetsConfigAPIRequest) GetAssetsConfigList() []AssetsConfig {
    return r._assetsConfigList
}
// MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveAssetsConfigAPIRequest) SetMiniAppId(_miniAppId string) error {
    r._miniAppId = _miniAppId
    r.Set("mini_app_id", _miniAppId)
    return nil
}

// MiniAppId Getter
func (r TaobaoJstInteractiveAssetsConfigAPIRequest) GetMiniAppId() string {
    return r._miniAppId
}
