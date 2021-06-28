package qianniu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qianniu"
)

/* 
任务元查询接口 
taobao.qianniu.taskmetas.get

任务元查询接口
*/
func TaobaoQianniuTaskmetasGet(clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskmetasGetRequest, session string) (*qianniu.TaobaoQianniuTaskmetasGetAPIResponse, error) {
    var resp qianniu.TaobaoQianniuTaskmetasGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
