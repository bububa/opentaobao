package idleparttime

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idleparttime"
)

/* 
兼职岗位同步 
alibaba.idle.parttime.jobsync

服务商同步岗位信息给闲鱼
*/
func AlibabaIdleParttimeJobsync(clt *core.SDKClient, req *idleparttime.AlibabaIdleParttimeJobsyncAPIRequest, session string) (*idleparttime.AlibabaIdleParttimeJobsyncAPIResponse, error) {
    var resp idleparttime.AlibabaIdleParttimeJobsyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
