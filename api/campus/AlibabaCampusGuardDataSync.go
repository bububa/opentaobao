package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
卡巴数据同步 
alibaba.campus.guard.data.sync

数据同步门禁系统
*/
func AlibabaCampusGuardDataSync(clt *core.SDKClient, req *campus.AlibabaCampusGuardDataSyncRequest, session string) (*campus.AlibabaCampusGuardDataSyncAPIResponse, error) {
    var resp campus.AlibabaCampusGuardDataSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
