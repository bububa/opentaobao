package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmallrefundget 获取OpenMall退款单详情
// taobao.openmall.refund.get
//
// 获取OpenMall退款单详情
func Taobaoopenmallrefundget(clt *core.SDKClient, req *openmall.TaobaoopenmallrefundgetAPIRequest, session string) (*openmall.TaobaoopenmallrefundgetAPIResponse, error) {
	var resp openmall.TaobaoopenmallrefundgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
