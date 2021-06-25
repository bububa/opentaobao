package qianniu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qianniu"
)

/* 
更新任务元数据 
taobao.qianniu.taskmeta.update

由任务发起者调用
*/
func TaobaoQianniuTaskmetaUpdate(clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskmetaUpdateRequest, session string) (*qianniu.TaobaoQianniuTaskmetaUpdateResponse, error) {
    var resp qianniu.TaobaoQianniuTaskmetaUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
