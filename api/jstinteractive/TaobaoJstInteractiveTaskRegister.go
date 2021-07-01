package jstinteractive

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jstinteractive"
)

/* 
互动任务开通接口 
taobao.jst.interactive.task.register

调用互动任务开通接口为小程序开通互动任务
*/
func TaobaoJstInteractiveTaskRegister(clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractiveTaskRegisterAPIRequest, session string) (*jstinteractive.TaobaoJstInteractiveTaskRegisterAPIResponse, error) {
    var resp jstinteractive.TaobaoJstInteractiveTaskRegisterAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
