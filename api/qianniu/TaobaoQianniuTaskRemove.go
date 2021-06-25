package qianniu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qianniu"
)

/* 
轻任务删除接口 
taobao.qianniu.task.remove

轻任务删除接口。
*/
func TaobaoQianniuTaskRemove(clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskRemoveRequest, session string) (*qianniu.TaobaoQianniuTaskRemoveResponse, error) {
    var resp qianniu.TaobaoQianniuTaskRemoveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
