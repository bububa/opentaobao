package qianniu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qianniu"
)

/* 
取消轻任务 
taobao.qianniu.task.cancel

由任务发起者调用
*/
func TaobaoQianniuTaskCancel(clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskCancelAPIRequest, session string) (*qianniu.TaobaoQianniuTaskCancelAPIResponse, error) {
    var resp qianniu.TaobaoQianniuTaskCancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
