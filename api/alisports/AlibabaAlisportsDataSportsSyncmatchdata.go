package alisports

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alisports"
)

/* 
阿里体育数据中心用户比赛数据同步接口 
alibaba.alisports.data.sports.syncmatchdata

阿里体育数据中心用户比赛数据同步接口
*/
func AlibabaAlisportsDataSportsSyncmatchdata(clt *core.SDKClient, req *alisports.AlibabaAlisportsDataSportsSyncmatchdataAPIRequest, session string) (*alisports.AlibabaAlisportsDataSportsSyncmatchdataAPIResponse, error) {
    var resp alisports.AlibabaAlisportsDataSportsSyncmatchdataAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
