package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询已配置的任务素材列表接口 APIRequest
taobao.jst.interactive.assets.configured.query

查询已配置任务素材列表
*/
type TaobaoJstInteractiveAssetsConfiguredQueryRequest struct {
    model.Params

    // 小程序id
    miniAppId   string 

}

func NewTaobaoJstInteractiveAssetsConfiguredQueryRequest() *TaobaoJstInteractiveAssetsConfiguredQueryRequest{
    return &TaobaoJstInteractiveAssetsConfiguredQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstInteractiveAssetsConfiguredQueryRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.assets.configured.query"
}

func (r TaobaoJstInteractiveAssetsConfiguredQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstInteractiveAssetsConfiguredQueryRequest) SetMiniAppId(miniAppId string) error {
    r.miniAppId = miniAppId
    r.Set("mini_app_id", miniAppId)
    return nil
}

func (r TaobaoJstInteractiveAssetsConfiguredQueryRequest) GetMiniAppId() string {
    return r.miniAppId
}

