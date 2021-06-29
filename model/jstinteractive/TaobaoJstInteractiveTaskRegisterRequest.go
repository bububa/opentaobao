package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动任务开通接口 API请求
taobao.jst.interactive.task.register

调用互动任务开通接口为小程序开通互动任务
*/
type TaobaoJstInteractiveTaskRegisterRequest struct {
    model.Params
    // 小程序id
    _miniAppId   string
}

// 初始化TaobaoJstInteractiveTaskRegisterRequest对象
func NewTaobaoJstInteractiveTaskRegisterRequest() *TaobaoJstInteractiveTaskRegisterRequest{
    return &TaobaoJstInteractiveTaskRegisterRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveTaskRegisterRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.task.register"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveTaskRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveTaskRegisterRequest) SetMiniAppId(_miniAppId string) error {
    r._miniAppId = _miniAppId
    r.Set("mini_app_id", _miniAppId)
    return nil
}

// MiniAppId Getter
func (r TaobaoJstInteractiveTaskRegisterRequest) GetMiniAppId() string {
    return r._miniAppId
}
