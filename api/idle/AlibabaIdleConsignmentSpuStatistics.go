package idle

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idle"
)

/* 
闲鱼帮卖同步服务商交易统计信息 
alibaba.idle.consignment.spu.statistics

闲鱼帮卖同步服务商交易统计信息
*/
func AlibabaIdleConsignmentSpuStatistics(clt *core.SDKClient, req *idle.AlibabaIdleConsignmentSpuStatisticsRequest, session string) (*idle.AlibabaIdleConsignmentSpuStatisticsAPIResponse, error) {
    var resp idle.AlibabaIdleConsignmentSpuStatisticsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
