package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcorpophotelbillsettlementquery 酒店结算记账查询接口
// alitrip.btrip.corpop.hotel.billsettlement.query
//
// 酒店结算记账查询接口
func Alitripbtripcorpophotelbillsettlementquery(clt *core.SDKClient, req *btrip.AlitripbtripcorpophotelbillsettlementqueryAPIRequest, session string) (*btrip.AlitripbtripcorpophotelbillsettlementqueryAPIResponse, error) {
	var resp btrip.AlitripbtripcorpophotelbillsettlementqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
