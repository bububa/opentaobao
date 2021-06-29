package alisports

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alisports"
)

/* 
阿里体育接入体脂秤数据 
alibaba.alisports.datacenter.datasync.fatscaledata

阿里体育数据中心接入体脂秤数据
*/
func AlibabaAlisportsDatacenterDatasyncFatscaledata(clt *core.SDKClient, req *alisports.AlibabaAlisportsDatacenterDatasyncFatscaledataRequest, session string) (*alisports.AlibabaAlisportsDatacenterDatasyncFatscaledataAPIResponse, error) {
    var resp alisports.AlibabaAlisportsDatacenterDatasyncFatscaledataAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
