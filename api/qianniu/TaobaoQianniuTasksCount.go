package qianniu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qianniu"
)

/* 
任务查询条数接口 
taobao.qianniu.tasks.count

任务查询条数接口
*/
func TaobaoQianniuTasksCount(clt *core.SDKClient, req *qianniu.TaobaoQianniuTasksCountRequest, session string) (*qianniu.TaobaoQianniuTasksCountAPIResponse, error) {
    var resp qianniu.TaobaoQianniuTasksCountAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
