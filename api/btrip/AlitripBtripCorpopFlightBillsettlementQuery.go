package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopFlightBillsettlementQuery 机票结算记账查询接口
// alitrip.btrip.corpop.flight.billsettlement.query
//
// 机票结算记账查询接口
func AlitripBtripCorpopFlightBillsettlementQuery(clt *core.SDKClient, req *btrip.AlitripBtripCorpopFlightBillsettlementQueryAPIRequest, session string) (*btrip.AlitripBtripCorpopFlightBillsettlementQueryAPIResponse, error) {
	var resp btrip.AlitripBtripCorpopFlightBillsettlementQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
