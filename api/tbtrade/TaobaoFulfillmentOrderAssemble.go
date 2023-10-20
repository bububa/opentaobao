package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// Taobaofulfillmentorderassemble 拆合单结果回传接口
// taobao.fulfillment.order.assemble
//
// 拆合单结果回传接口
func Taobaofulfillmentorderassemble(clt *core.SDKClient, req *tbtrade.TaobaofulfillmentorderassembleAPIRequest, session string) (*tbtrade.TaobaofulfillmentorderassembleAPIResponse, error) {
	var resp tbtrade.TaobaofulfillmentorderassembleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
