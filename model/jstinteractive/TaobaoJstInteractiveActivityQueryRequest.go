package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动任务活动查询接口 API请求
taobao.jst.interactive.activity.query

互动任务活动查询接口
*/
type TaobaoJstInteractiveActivityQueryRequest struct {
    model.Params
    // 小程序id
    _miniAppId   string
}

// 初始化TaobaoJstInteractiveActivityQueryRequest对象
func NewTaobaoJstInteractiveActivityQueryRequest() *TaobaoJstInteractiveActivityQueryRequest{
    return &TaobaoJstInteractiveActivityQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveActivityQueryRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.activity.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveActivityQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveActivityQueryRequest) SetMiniAppId(_miniAppId string) error {
    r._miniAppId = _miniAppId
    r.Set("mini_app_id", _miniAppId)
    return nil
}

// MiniAppId Getter
func (r TaobaoJstInteractiveActivityQueryRequest) GetMiniAppId() string {
    return r._miniAppId
}
