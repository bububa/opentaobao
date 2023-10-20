package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcorpopflightbillsettlementquery 机票结算记账查询接口
// alitrip.btrip.corpop.flight.billsettlement.query
//
// 机票结算记账查询接口
func Alitripbtripcorpopflightbillsettlementquery(clt *core.SDKClient, req *btrip.AlitripbtripcorpopflightbillsettlementqueryAPIRequest, session string) (*btrip.AlitripbtripcorpopflightbillsettlementqueryAPIResponse, error) {
	var resp btrip.AlitripbtripcorpopflightbillsettlementqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
