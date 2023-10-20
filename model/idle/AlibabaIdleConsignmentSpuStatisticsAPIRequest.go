package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleConsignmentSpuStatisticsAPIRequest 闲鱼帮卖同步服务商交易统计信息 API请求
// alibaba.idle.consignment.spu.statistics
//
// 闲鱼帮卖同步服务商交易统计信息
type AlibabaIdleConsignmentSpuStatisticsAPIRequest struct {
	model.Params
	// 入参
	_param *SpuStatistics
}

// NewAlibabaIdleConsignmentSpuStatisticsRequest 初始化AlibabaIdleConsignmentSpuStatisticsAPIRequest对象
func NewAlibabaIdleConsignmentSpuStatisticsRequest() *AlibabaIdleConsignmentSpuStatisticsAPIRequest {
	return &AlibabaIdleConsignmentSpuStatisticsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleConsignmentSpuStatisticsAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.consignment.spu.statistics"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleConsignmentSpuStatisticsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleConsignmentSpuStatisticsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaIdleConsignmentSpuStatisticsAPIRequest) SetParam(_param *SpuStatistics) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaIdleConsignmentSpuStatisticsAPIRequest) GetParam() *SpuStatistics {
	return r._param
}
