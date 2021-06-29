package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动任务列表查询接口 API请求
taobao.jst.interactive.task.query

查询用户的互动任务列表
*/
type TaobaoJstInteractiveTaskQueryRequest struct {
    model.Params
}

// 初始化TaobaoJstInteractiveTaskQueryRequest对象
func NewTaobaoJstInteractiveTaskQueryRequest() *TaobaoJstInteractiveTaskQueryRequest{
    return &TaobaoJstInteractiveTaskQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveTaskQueryRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.task.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveTaskQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
