package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼帮卖同步服务商交易统计信息 API请求
alibaba.idle.consignment.spu.statistics

闲鱼帮卖同步服务商交易统计信息
*/
type AlibabaIdleConsignmentSpuStatisticsRequest struct {
    model.Params
    // 入参
    param   *SpuStatistics
}

// 初始化AlibabaIdleConsignmentSpuStatisticsRequest对象
func NewAlibabaIdleConsignmentSpuStatisticsRequest() *AlibabaIdleConsignmentSpuStatisticsRequest{
    return &AlibabaIdleConsignmentSpuStatisticsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleConsignmentSpuStatisticsRequest) GetApiMethodName() string {
    return "alibaba.idle.consignment.spu.statistics"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleConsignmentSpuStatisticsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaIdleConsignmentSpuStatisticsRequest) SetParam(param *SpuStatistics) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaIdleConsignmentSpuStatisticsRequest) GetParam() *SpuStatistics {
    return r.param
}
