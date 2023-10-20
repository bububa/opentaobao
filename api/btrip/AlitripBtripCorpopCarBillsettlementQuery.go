package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopCarBillsettlementQuery 用车结算记账查询接口
// alitrip.btrip.corpop.car.billsettlement.query
//
// 用车结算记账查询接口
func AlitripBtripCorpopCarBillsettlementQuery(clt *core.SDKClient, req *btrip.AlitripBtripCorpopCarBillsettlementQueryAPIRequest, resp *btrip.AlitripBtripCorpopCarBillsettlementQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
