package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleConsignmentSpuStatisticsAPIRequest
闲鱼帮卖同步服务商交易统计信息 API请求
alibaba.idle.consignment.spu.statistics

闲鱼帮卖同步服务商交易统计信息 */
type AlibabaIdleConsignmentSpuStatisticsAPIRequest struct {
	model.Params
	// 入参
	_param *SpuStatistics
}

// New
