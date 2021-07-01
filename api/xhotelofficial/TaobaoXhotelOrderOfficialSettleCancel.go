package xhotelofficial

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelofficial"
)

/* TaobaoXhotelOrderOfficialSettleCancel
官网信用住取消结账
taobao.xhotel.order.official.settle.cancel

用于官网信用住取消结账 */
func TaobaoXhotelOrderOfficialSettleCancel(clt *core.SDKClient, req *xhotelofficial.TaobaoXhotelOrderOfficialSettleCancelAPIRequest, session string) (*xhotelofficial.TaobaoXhotelOrderOfficialSettleCancelAPIResponse, error) {
	var resp xhotelofficial.TaobaoXhotelOrderOfficialSettleCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
