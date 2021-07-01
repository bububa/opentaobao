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
type AlibabaIdleConsignmentSpuStatisticsAPIRequest struct {
    model.Params
    // 入参
    _param   *SpuStatistics
}

// 初始化AlibabaIdleConsignmentSpuStatisticsAPIRequest对象
func NewAlibabaIdleConsignmentSpuStatisticsRequest() *AlibabaIdleConsignmentSpuStatisticsAPIRequest{
    return &AlibabaIdleConsignmentSpuStatisticsAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleConsignmentSpuStatisticsAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.consignment.spu.statistics"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleConsignmentSpuStatisticsAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaIdleConsignmentSpuStatisticsAPIRequest) SetParam(_param *SpuStatistics) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaIdleConsignmentSpuStatisticsAPIRequest) GetParam() *SpuStatistics {
    return r._param
}
