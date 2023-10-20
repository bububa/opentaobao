package deliveryvoucher

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

// Taobaogamedeliveryvouchersendvoucher 提货券发券接口
// taobao.game.deliveryvoucher.sendvoucher
//
// 提货券发券接口：同步券和订单的关联信息
func Taobaogamedeliveryvouchersendvoucher(clt *core.SDKClient, req *deliveryvoucher.TaobaogamedeliveryvouchersendvoucherAPIRequest, session string) (*deliveryvoucher.TaobaogamedeliveryvouchersendvoucherAPIResponse, error) {
	var resp deliveryvoucher.TaobaogamedeliveryvouchersendvoucherAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
