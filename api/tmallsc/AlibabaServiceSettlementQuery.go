package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaServiceSettlementQuery 服务平台结算单明细查询服务
// alibaba.service.settlement.query
//
// 给服务商提供结算单明细查询功能
func AlibabaServiceSettlementQuery(clt *core.SDKClient, req *tmallsc.AlibabaServiceSettlementQueryAPIRequest, resp *tmallsc.AlibabaServiceSettlementQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
