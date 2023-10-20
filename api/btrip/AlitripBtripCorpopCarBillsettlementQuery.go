package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcorpopcarbillsettlementquery 用车结算记账查询接口
// alitrip.btrip.corpop.car.billsettlement.query
//
// 用车结算记账查询接口
func Alitripbtripcorpopcarbillsettlementquery(clt *core.SDKClient, req *btrip.AlitripbtripcorpopcarbillsettlementqueryAPIRequest, session string) (*btrip.AlitripbtripcorpopcarbillsettlementqueryAPIResponse, error) {
	var resp btrip.AlitripbtripcorpopcarbillsettlementqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
