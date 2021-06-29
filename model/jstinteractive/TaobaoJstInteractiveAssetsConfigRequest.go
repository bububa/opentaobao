package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
任务素材配置接口 APIRequest
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

func NewTaobaoJstInteractiveAssetsConfigRequest() *TaobaoJstInteractiveAssetsConfigRequest{
    return &TaobaoJstInteractiveAssetsConfigRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstInteractiveAssetsConfigRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.assets.config"
}

func (r TaobaoJstInteractiveAssetsConfigRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstInteractiveAssetsConfigRequest) SetAssetsConfigList(assetsConfigList []AssetsConfig) error {
    r.assetsConfigList = assetsConfigList
    r.Set("assets_config_list", assetsConfigList)
    return nil
}

func (r TaobaoJstInteractiveAssetsConfigRequest) GetAssetsConfigList() []AssetsConfig {
    return r.assetsConfigList
}

func (r *TaobaoJstInteractiveAssetsConfigRequest) SetMiniAppId(miniAppId string) error {
    r.miniAppId = miniAppId
    r.Set("mini_app_id", miniAppId)
    return nil
}

func (r TaobaoJstInteractiveAssetsConfigRequest) GetMiniAppId() string {
    return r.miniAppId
}

