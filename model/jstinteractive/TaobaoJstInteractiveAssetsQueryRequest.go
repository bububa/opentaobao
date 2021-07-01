package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询可配置任务素材接口 API请求
taobao.jst.interactive.assets.query

查询可配置任务素材列表，用以配置任务素材
*/
type TaobaoJstInteractiveAssetsQueryAPIRequest struct {
    model.Params
    // 小程序id
    _miniAppId   string
}

// 初始化TaobaoJstInteractiveAssetsQueryAPIRequest对象
func NewTaobaoJstInteractiveAssetsQueryRequest() *TaobaoJstInteractiveAssetsQueryAPIRequest{
    return &TaobaoJstInteractiveAssetsQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveAssetsQueryAPIRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.assets.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveAssetsQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveAssetsQueryAPIRequest) SetMiniAppId(_miniAppId string) error {
    r._miniAppId = _miniAppId
    r.Set("mini_app_id", _miniAppId)
    return nil
}

// MiniAppId Getter
func (r TaobaoJstInteractiveAssetsQueryAPIRequest) GetMiniAppId() string {
    return r._miniAppId
}
