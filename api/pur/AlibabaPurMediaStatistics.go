package pur

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/pur"
)

/* 
新媒体统计信息 
alibaba.pur.media.statistics

清博同步新媒体的统计信息给到采购平台
*/
func AlibabaPurMediaStatistics(clt *core.SDKClient, req *pur.AlibabaPurMediaStatisticsRequest, session string) (*pur.AlibabaPurMediaStatisticsAPIResponse, error) {
    var resp pur.AlibabaPurMediaStatisticsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
