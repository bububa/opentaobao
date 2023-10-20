package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Cainiaorefundrefundactionsdisplay 退货退款操作的展示信息(展现给买家)
// cainiao.refund.refundactions.display
//
// 退货退款操作的展示信息(展现给买家)
func Cainiaorefundrefundactionsdisplay(clt *core.SDKClient, req *trade.CainiaorefundrefundactionsdisplayAPIRequest, session string) (*trade.CainiaorefundrefundactionsdisplayAPIResponse, error) {
	var resp trade.CainiaorefundrefundactionsdisplayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
