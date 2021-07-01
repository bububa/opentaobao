package deliveryvoucher

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

/* TaobaoGameDeliveryvoucherSendvoucher
提货券发券接口
taobao.game.deliveryvoucher.sendvoucher

提货券发券接口：同步券和订单的关联信息 */
func TaobaoGameDeliveryvoucherSendvoucher(clt *core.SDKClient, req *deliveryvoucher.TaobaoGameDeliveryvoucherSendvoucherAPIRequest, session string) (*deliveryvoucher.TaobaoGameDeliveryvoucherSendvoucherAPIResponse, error) {
	var resp deliveryvoucher.TaobaoGameDeliveryvoucherSendvoucherAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
