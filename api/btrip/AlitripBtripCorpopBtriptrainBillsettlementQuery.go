package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopBtriptrainBillsettlementQuery 商旅火车票结算记账查询接口
// alitrip.btrip.corpop.btriptrain.billsettlement.query
//
// 商旅火车票结算记账查询接口
func AlitripBtripCorpopBtriptrainBillsettlementQuery(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest, resp *btrip.AlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
