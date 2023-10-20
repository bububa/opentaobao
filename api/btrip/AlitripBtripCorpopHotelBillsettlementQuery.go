package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopHotelBillsettlementQuery 酒店结算记账查询接口
// alitrip.btrip.corpop.hotel.billsettlement.query
//
// 酒店结算记账查询接口
func AlitripBtripCorpopHotelBillsettlementQuery(clt *core.SDKClient, req *btrip.AlitripBtripCorpopHotelBillsettlementQueryAPIRequest, session string) (*btrip.AlitripBtripCorpopHotelBillsettlementQueryAPIResponse, error) {
	var resp btrip.AlitripBtripCorpopHotelBillsettlementQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
