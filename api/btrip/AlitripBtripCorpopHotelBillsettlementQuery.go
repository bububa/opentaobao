package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopHotelBillsettlementQuery 酒店结算记账查询接口
// alitrip.btrip.corpop.hotel.billsettlement.query
//
// 酒店结算记账查询接口
func AlitripBtripCorpopHotelBillsettlementQuery(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCorpopHotelBillsettlementQueryAPIRequest, resp *btrip.AlitripBtripCorpopHotelBillsettlementQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
