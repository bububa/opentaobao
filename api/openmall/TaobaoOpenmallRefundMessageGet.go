package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

/* TaobaoOpenmallRefundMessageGet
openmall获取退款单留言
taobao.openmall.refund.message.get

openmall获取退款单留言 */
func TaobaoOpenmallRefundMessageGet(clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundMessageGetAPIRequest, session string) (*openmall.TaobaoOpenmallRefundMessageGetAPIResponse, error) {
	var resp openmall.TaobaoOpenmallRefundMessageGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
