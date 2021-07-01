package idleparttime

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idleparttime"
)

/* 
兼职通知接口 
alibaba.idle.parttime.notify

服务商侧有岗位状态变更对我们进行通知(岗位关闭, 录取状态)
*/
func AlibabaIdleParttimeNotify(clt *core.SDKClient, req *idleparttime.AlibabaIdleParttimeNotifyAPIRequest, session string) (*idleparttime.AlibabaIdleParttimeNotifyAPIResponse, error) {
    var resp idleparttime.AlibabaIdleParttimeNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
