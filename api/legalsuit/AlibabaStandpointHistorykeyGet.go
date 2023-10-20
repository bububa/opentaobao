package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaStandpointHistorykeyGet 查询历史数据
// alibaba.standpoint.historykey.get
//
// 查询历史数据
func AlibabaStandpointHistorykeyGet(clt *core.SDKClient, req *legalsuit.AlibabaStandpointHistorykeyGetAPIRequest, resp *legalsuit.AlibabaStandpointHistorykeyGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
