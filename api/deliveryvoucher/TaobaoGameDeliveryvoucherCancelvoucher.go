package deliveryvoucher

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

// Taobaogamedeliveryvouchercancelvoucher 作废券
// taobao.game.deliveryvoucher.cancelvoucher
//
// 提货券发券接口：同步券和订单的关联信息
func Taobaogamedeliveryvouchercancelvoucher(clt *core.SDKClient, req *deliveryvoucher.TaobaogamedeliveryvouchercancelvoucherAPIRequest, session string) (*deliveryvoucher.TaobaogamedeliveryvouchercancelvoucherAPIResponse, error) {
	var resp deliveryvoucher.TaobaogamedeliveryvouchercancelvoucherAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
