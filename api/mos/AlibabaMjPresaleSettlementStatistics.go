package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

/* AlibabaMjPresaleSettlementStatistics
预购结算数据统计
alibaba.mj.presale.settlement.statistics

预购结算数据统计 */
func AlibabaMjPresaleSettlementStatistics(clt *core.SDKClient, req *mos.AlibabaMjPresaleSettlementStatisticsAPIRequest, session string) (*mos.AlibabaMjPresaleSettlementStatisticsAPIResponse, error) {
	var resp mos.AlibabaMjPresaleSettlementStatisticsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
