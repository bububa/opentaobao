package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动任务列表查询接口 APIRequest
taobao.jst.interactive.task.query

查询用户的互动任务列表
*/
type TaobaoJstInteractiveTaskQueryRequest struct {
    model.Params

}

func NewTaobaoJstInteractiveTaskQueryRequest() *TaobaoJstInteractiveTaskQueryRequest{
    return &TaobaoJstInteractiveTaskQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstInteractiveTaskQueryRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.task.query"
}

func (r TaobaoJstInteractiveTaskQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


