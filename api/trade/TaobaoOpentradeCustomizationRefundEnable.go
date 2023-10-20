package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Taobaoopentradecustomizationrefundenable 定制订单设置允许仅退款
// taobao.opentrade.customization.refund.enable
//
// 定制订单设置允许仅退款
func Taobaoopentradecustomizationrefundenable(clt *core.SDKClient, req *trade.TaobaoopentradecustomizationrefundenableAPIRequest, session string) (*trade.TaobaoopentradecustomizationrefundenableAPIResponse, error) {
	var resp trade.TaobaoopentradecustomizationrefundenableAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
