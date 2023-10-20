package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurMediaStatistics 新媒体统计信息
// alibaba.pur.media.statistics
//
// 清博同步新媒体的统计信息给到采购平台
func AlibabaPurMediaStatistics(clt *core.SDKClient, req *pur.AlibabaPurMediaStatisticsAPIRequest, resp *pur.AlibabaPurMediaStatisticsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
