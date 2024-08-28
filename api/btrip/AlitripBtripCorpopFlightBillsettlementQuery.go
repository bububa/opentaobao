package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopFlightBillsettlementQuery 机票结算记账查询接口
// alitrip.btrip.corpop.flight.billsettlement.query
//
// 机票结算记账查询接口
func AlitripBtripCorpopFlightBillsettlementQuery(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCorpopFlightBillsettlementQueryAPIRequest, resp *btrip.AlitripBtripCorpopFlightBillsettlementQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
