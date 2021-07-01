package deliveryvoucher

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

/* TaobaoGameDeliveryvoucherCancelvoucher
作废券
taobao.game.deliveryvoucher.cancelvoucher

提货券发券接口：同步券和订单的关联信息 */
func TaobaoGameDeliveryvoucherCancelvoucher(clt *core.SDKClient, req *deliveryvoucher.TaobaoGameDeliveryvoucherCancelvoucherAPIRequest, session string) (*deliveryvoucher.TaobaoGameDeliveryvoucherCancelvoucherAPIResponse, error) {
	var resp deliveryvoucher.TaobaoGameDeliveryvoucherCancelvoucherAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
