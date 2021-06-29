package alisports

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alisports"
)

/* 
阿里体育数据中心用户当天累积数据同步接口 
alibaba.alisports.data.sports.syncstatdata

阿里体育数据中心用户当天累积数据同步接口
*/
func AlibabaAlisportsDataSportsSyncstatdata(clt *core.SDKClient, req *alisports.AlibabaAlisportsDataSportsSyncstatdataRequest, session string) (*alisports.AlibabaAlisportsDataSportsSyncstatdataAPIResponse, error) {
    var resp alisports.AlibabaAlisportsDataSportsSyncstatdataAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
