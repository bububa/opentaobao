package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmallrefundmessageget openmall获取退款单留言
// taobao.openmall.refund.message.get
//
// openmall获取退款单留言
func Taobaoopenmallrefundmessageget(clt *core.SDKClient, req *openmall.TaobaoopenmallrefundmessagegetAPIRequest, session string) (*openmall.TaobaoopenmallrefundmessagegetAPIResponse, error) {
	var resp openmall.TaobaoopenmallrefundmessagegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
