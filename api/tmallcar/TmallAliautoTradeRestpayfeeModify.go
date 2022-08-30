package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoTradeRestpayfeeModify 分阶段订单尾款改价
// tmall.aliauto.trade.restpayfee.modify
//
// 汽车商家通过此api修改整车分阶段订单的尾款金额
func TmallAliautoTradeRestpayfeeModify(clt *core.SDKClient, req *tmallcar.TmallAliautoTradeRestpayfeeModifyAPIRequest, session string) (*tmallcar.TmallAliautoTradeRestpayfeeModifyAPIResponse, error) {
	var resp tmallcar.TmallAliautoTradeRestpayfeeModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
