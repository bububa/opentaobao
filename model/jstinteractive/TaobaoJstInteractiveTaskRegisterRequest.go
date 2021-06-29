package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动任务开通接口 APIRequest
taobao.jst.interactive.task.register

调用互动任务开通接口为小程序开通互动任务
*/
type TaobaoJstInteractiveTaskRegisterRequest struct {
    model.Params

    // 小程序id
    miniAppId   string 

}

func NewTaobaoJstInteractiveTaskRegisterRequest() *TaobaoJstInteractiveTaskRegisterRequest{
    return &TaobaoJstInteractiveTaskRegisterRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstInteractiveTaskRegisterRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.task.register"
}

func (r TaobaoJstInteractiveTaskRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstInteractiveTaskRegisterRequest) SetMiniAppId(miniAppId string) error {
    r.miniAppId = miniAppId
    r.Set("mini_app_id", miniAppId)
    return nil
}

func (r TaobaoJstInteractiveTaskRegisterRequest) GetMiniAppId() string {
    return r.miniAppId
}

