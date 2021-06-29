package idleparttime

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idleparttime"
)

/* 
兼职同步日志 
alibaba.idle.parttime.synclog

提供给供应商查询的接口
*/
func AlibabaIdleParttimeSynclog(clt *core.SDKClient, req *idleparttime.AlibabaIdleParttimeSynclogRequest, session string) (*idleparttime.AlibabaIdleParttimeSynclogAPIResponse, error) {
    var resp idleparttime.AlibabaIdleParttimeSynclogAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
