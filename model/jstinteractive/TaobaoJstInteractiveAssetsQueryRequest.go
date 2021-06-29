package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询可配置任务素材接口 APIRequest
taobao.jst.interactive.assets.query

查询可配置任务素材列表，用以配置任务素材
*/
type TaobaoJstInteractiveAssetsQueryRequest struct {
    model.Params

    // 小程序id
    miniAppId   string 

}

func NewTaobaoJstInteractiveAssetsQueryRequest() *TaobaoJstInteractiveAssetsQueryRequest{
    return &TaobaoJstInteractiveAssetsQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstInteractiveAssetsQueryRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.assets.query"
}

func (r TaobaoJstInteractiveAssetsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstInteractiveAssetsQueryRequest) SetMiniAppId(miniAppId string) error {
    r.miniAppId = miniAppId
    r.Set("mini_app_id", miniAppId)
    return nil
}

func (r TaobaoJstInteractiveAssetsQueryRequest) GetMiniAppId() string {
    return r.miniAppId
}

