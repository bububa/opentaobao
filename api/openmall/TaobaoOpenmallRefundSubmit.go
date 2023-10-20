package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmallrefundsubmit 提交OpenMall退款单物流
// taobao.openmall.refund.submit
//
// 提交OpenMall退款单物流
func Taobaoopenmallrefundsubmit(clt *core.SDKClient, req *openmall.TaobaoopenmallrefundsubmitAPIRequest, session string) (*openmall.TaobaoopenmallrefundsubmitAPIResponse, error) {
	var resp openmall.TaobaoopenmallrefundsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
