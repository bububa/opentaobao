package rhino

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rhino"
)

// Taobaorhinosupplychainoutboundpickingcomplete 【WMS005】接收成衣捡配完成通知
// taobao.rhino.supplychain.outbound.pickingcomplete
//
// 接收成衣捡配完成通知,WMS005
func Taobaorhinosupplychainoutboundpickingcomplete(clt *core.SDKClient, req *rhino.TaobaorhinosupplychainoutboundpickingcompleteAPIRequest, session string) (*rhino.TaobaorhinosupplychainoutboundpickingcompleteAPIResponse, error) {
	var resp rhino.TaobaorhinosupplychainoutboundpickingcompleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
