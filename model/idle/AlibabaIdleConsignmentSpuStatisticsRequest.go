package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼帮卖同步服务商交易统计信息 APIRequest
alibaba.idle.consignment.spu.statistics

闲鱼帮卖同步服务商交易统计信息
*/
type AlibabaIdleConsignmentSpuStatisticsRequest struct {
    model.Params

    // 入参
    param   *SpuStatistics 

}

func NewAlibabaIdleConsignmentSpuStatisticsRequest() *AlibabaIdleConsignmentSpuStatisticsRequest{
    return &AlibabaIdleConsignmentSpuStatisticsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleConsignmentSpuStatisticsRequest) GetApiMethodName() string {
    return "alibaba.idle.consignment.spu.statistics"
}

func (r AlibabaIdleConsignmentSpuStatisticsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleConsignmentSpuStatisticsRequest) SetParam(param *SpuStatistics) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaIdleConsignmentSpuStatisticsRequest) GetParam() *SpuStatistics {
    return r.param
}

