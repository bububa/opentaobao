package alisports

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alisports"
)

/* 
阿里体育同步跑步机设备数据 
alibaba.alisports.datacenter.datasync.treadmill

合作方向阿里体育同步跑步机设备的数据
*/
func AlibabaAlisportsDatacenterDatasyncTreadmill(clt *core.SDKClient, req *alisports.AlibabaAlisportsDatacenterDatasyncTreadmillRequest, session string) (*alisports.AlibabaAlisportsDatacenterDatasyncTreadmillAPIResponse, error) {
    var resp alisports.AlibabaAlisportsDatacenterDatasyncTreadmillAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
