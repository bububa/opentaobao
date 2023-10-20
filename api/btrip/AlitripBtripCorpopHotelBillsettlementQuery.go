package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopHotelBillsettlementQuery 酒店结算记账查询接口
// alitrip.btrip.corpop.hotel.billsettlement.query
//
// 酒店结算记账查询接口
func AlitripBtripCorpopHotelBillsettlementQuery(clt *core.SDKClient, req *btrip.AlitripBtripCorpopHotelBillsettlementQueryAPIRequest, resp *btrip.AlitripBtripCorpopHotelBillsettlementQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
