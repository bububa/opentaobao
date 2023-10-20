package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcorpopbtriptrainbillsettlementquery 商旅火车票结算记账查询接口
// alitrip.btrip.corpop.btriptrain.billsettlement.query
//
// 商旅火车票结算记账查询接口
func Alitripbtripcorpopbtriptrainbillsettlementquery(clt *core.SDKClient, req *btrip.AlitripbtripcorpopbtriptrainbillsettlementqueryAPIRequest, session string) (*btrip.AlitripbtripcorpopbtriptrainbillsettlementqueryAPIResponse, error) {
	var resp btrip.AlitripbtripcorpopbtriptrainbillsettlementqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
