package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopCarBillsettlementQuery 用车结算记账查询接口
// alitrip.btrip.corpop.car.billsettlement.query
//
// 用车结算记账查询接口
func AlitripBtripCorpopCarBillsettlementQuery(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCorpopCarBillsettlementQueryAPIRequest, resp *btrip.AlitripBtripCorpopCarBillsettlementQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
