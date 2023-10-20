package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobustvmpayorderset 自助机条形码被动支付
// taobao.bus.tvmpayorder.set
//
// 汽车票线下自助机条形码支付
func Taobaobustvmpayorderset(clt *core.SDKClient, req *bus.TaobaobustvmpayordersetAPIRequest, session string) (*bus.TaobaobustvmpayordersetAPIResponse, error) {
	var resp bus.TaobaobustvmpayordersetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
