package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleConsignmentSpuStatistics 闲鱼帮卖同步服务商交易统计信息
// alibaba.idle.consignment.spu.statistics
//
// 闲鱼帮卖同步服务商交易统计信息
func AlibabaIdleConsignmentSpuStatistics(clt *core.SDKClient, req *idle.AlibabaIdleConsignmentSpuStatisticsAPIRequest, resp *idle.AlibabaIdleConsignmentSpuStatisticsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
