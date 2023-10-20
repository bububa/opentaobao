package deliveryvoucher

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

// Taobaogamedeliveryvoucherrollbackvoucher 回滚券
// taobao.game.deliveryvoucher.rollbackvoucher
//
// 提货券发券接口：同步券和订单的关联信息
func Taobaogamedeliveryvoucherrollbackvoucher(clt *core.SDKClient, req *deliveryvoucher.TaobaogamedeliveryvoucherrollbackvoucherAPIRequest, session string) (*deliveryvoucher.TaobaogamedeliveryvoucherrollbackvoucherAPIResponse, error) {
	var resp deliveryvoucher.TaobaogamedeliveryvoucherrollbackvoucherAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
