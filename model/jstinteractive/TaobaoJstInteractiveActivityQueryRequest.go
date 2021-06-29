package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动任务活动查询接口 APIRequest
taobao.jst.interactive.activity.query

互动任务活动查询接口
*/
type TaobaoJstInteractiveActivityQueryRequest struct {
    model.Params

    // 小程序id
    miniAppId   string 

}

func NewTaobaoJstInteractiveActivityQueryRequest() *TaobaoJstInteractiveActivityQueryRequest{
    return &TaobaoJstInteractiveActivityQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstInteractiveActivityQueryRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.activity.query"
}

func (r TaobaoJstInteractiveActivityQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstInteractiveActivityQueryRequest) SetMiniAppId(miniAppId string) error {
    r.miniAppId = miniAppId
    r.Set("mini_app_id", miniAppId)
    return nil
}

func (r TaobaoJstInteractiveActivityQueryRequest) GetMiniAppId() string {
    return r.miniAppId
}

