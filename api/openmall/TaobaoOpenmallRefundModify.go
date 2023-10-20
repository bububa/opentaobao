package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmallrefundmodify 修改OpenMall退款申请
// taobao.openmall.refund.modify
//
// 修改OpenMall退款申请
func Taobaoopenmallrefundmodify(clt *core.SDKClient, req *openmall.TaobaoopenmallrefundmodifyAPIRequest, session string) (*openmall.TaobaoopenmallrefundmodifyAPIResponse, error) {
	var resp openmall.TaobaoopenmallrefundmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
