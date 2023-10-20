package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmallrefundclose 关闭OpenMall退款单
// taobao.openmall.refund.close
//
// 关闭OpenMall退款单
func Taobaoopenmallrefundclose(clt *core.SDKClient, req *openmall.TaobaoopenmallrefundcloseAPIRequest, session string) (*openmall.TaobaoopenmallrefundcloseAPIResponse, error) {
	var resp openmall.TaobaoopenmallrefundcloseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
