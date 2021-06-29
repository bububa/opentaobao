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
type TaobaoJstInteractiveAssetsConfigRequest struct {
    model.Params
    // []
    assetsConfigList   []AssetsConfig
    // 小程序id
    miniAppId   string
}

// 初始化TaobaoJstInteractiveAssetsConfigRequest对象
func NewTaobaoJstInteractiveAssetsConfigRequest() *TaobaoJstInteractiveAssetsConfigRequest{
    return &TaobaoJstInteractiveAssetsConfigRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveAssetsConfigRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.assets.config"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveAssetsConfigRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AssetsConfigList Setter
// []
func (r *TaobaoJstInteractiveAssetsConfigRequest) SetAssetsConfigList(assetsConfigList []AssetsConfig) error {
    r.assetsConfigList = assetsConfigList
    r.Set("assets_config_list", assetsConfigList)
    return nil
}

// AssetsConfigList Getter
func (r TaobaoJstInteractiveAssetsConfigRequest) GetAssetsConfigList() []AssetsConfig {
    return r.assetsConfigList
}
// MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveAssetsConfigRequest) SetMiniAppId(miniAppId string) error {
    r.miniAppId = miniAppId
    r.Set("mini_app_id", miniAppId)
    return nil
}

// MiniAppId Getter
func (r TaobaoJstInteractiveAssetsConfigRequest) GetMiniAppId() string {
    return r.miniAppId
}
