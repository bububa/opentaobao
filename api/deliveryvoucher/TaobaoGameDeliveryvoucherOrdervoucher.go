package deliveryvoucher

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

// Taobaogamedeliveryvoucherordervoucher 预约接口
// taobao.game.deliveryvoucher.ordervoucher
//
// 提货券发券接口：同步券和订单的关联信息
func Taobaogamedeliveryvoucherordervoucher(clt *core.SDKClient, req *deliveryvoucher.TaobaogamedeliveryvoucherordervoucherAPIRequest, session string) (*deliveryvoucher.TaobaogamedeliveryvoucherordervoucherAPIResponse, error) {
	var resp deliveryvoucher.TaobaogamedeliveryvoucherordervoucherAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
