package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopBtriptrainBillsettlementQuery 商旅火车票结算记账查询接口
// alitrip.btrip.corpop.btriptrain.billsettlement.query
//
// 商旅火车票结算记账查询接口
func AlitripBtripCorpopBtriptrainBillsettlementQuery(clt *core.SDKClient, req *btrip.AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest, resp *btrip.AlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
