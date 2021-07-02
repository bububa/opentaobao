package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundGet 获取OpenMall退款单详情
// taobao.openmall.refund.get
//
// 获取OpenMall退款单详情
func TaobaoOpenmallRefundGet(clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundGetAPIRequest, session string) (*openmall.TaobaoOpenmallRefundGetAPIResponse, error) {
	var resp openmall.TaobaoOpenmallRefundGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
