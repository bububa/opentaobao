package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaStandpointHistorykeyGet 查询历史数据
// alibaba.standpoint.historykey.get
//
// 查询历史数据
func AlibabaStandpointHistorykeyGet(clt *core.SDKClient, req *legalsuit.AlibabaStandpointHistorykeyGetAPIRequest, session string) (*legalsuit.AlibabaStandpointHistorykeyGetAPIResponse, error) {
	var resp legalsuit.AlibabaStandpointHistorykeyGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
