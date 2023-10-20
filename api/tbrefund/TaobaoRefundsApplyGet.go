package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Taobaorefundsapplyget 查询买家申请的退款列表
// taobao.refunds.apply.get
//
// 查询买家申请的退款列表，且查询外店的退款列表时需要指定交易类型
func Taobaorefundsapplyget(clt *core.SDKClient, req *tbrefund.TaobaorefundsapplygetAPIRequest, session string) (*tbrefund.TaobaorefundsapplygetAPIResponse, error) {
	var resp tbrefund.TaobaorefundsapplygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
