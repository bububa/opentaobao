package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoRdcAligeniusRefundsCheck 退款信息审核
// taobao.rdc.aligenius.refunds.check
//
// 根据退款信息，对退款单进行审核
func TaobaoRdcAligeniusRefundsCheck(clt *core.SDKClient, req *util.TaobaoRdcAligeniusRefundsCheckAPIRequest, session string) (*util.TaobaoRdcAligeniusRefundsCheckAPIResponse, error) {
	var resp util.TaobaoRdcAligeniusRefundsCheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
