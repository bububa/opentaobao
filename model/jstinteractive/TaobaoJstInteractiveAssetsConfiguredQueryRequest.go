package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询已配置的任务素材列表接口 API请求
taobao.jst.interactive.assets.configured.query

查询已配置任务素材列表
*/
type TaobaoJstInteractiveAssetsConfiguredQueryRequest struct {
    model.Params
    // 小程序id
    miniAppId   string
}

// 初始化TaobaoJstInteractiveAssetsConfiguredQueryRequest对象
func NewTaobaoJstInteractiveAssetsConfiguredQueryRequest() *TaobaoJstInteractiveAssetsConfiguredQueryRequest{
    return &TaobaoJstInteractiveAssetsConfiguredQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveAssetsConfiguredQueryRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.assets.configured.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveAssetsConfiguredQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveAssetsConfiguredQueryRequest) SetMiniAppId(miniAppId string) error {
    r.miniAppId = miniAppId
    r.Set("mini_app_id", miniAppId)
    return nil
}

// MiniAppId Getter
func (r TaobaoJstInteractiveAssetsConfiguredQueryRequest) GetMiniAppId() string {
    return r.miniAppId
}
