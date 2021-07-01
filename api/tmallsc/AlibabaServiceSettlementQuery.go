package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

/* AlibabaServiceSettlementQuery
服务平台结算单明细查询服务
alibaba.service.settlement.query

给服务商提供结算单明细查询功能 */
func AlibabaServiceSettlementQuery(clt *core.SDKClient, req *tmallsc.AlibabaServiceSettlementQueryAPIRequest, session string) (*tmallsc.AlibabaServiceSettlementQueryAPIResponse, error) {
	var resp tmallsc.AlibabaServiceSettlementQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
