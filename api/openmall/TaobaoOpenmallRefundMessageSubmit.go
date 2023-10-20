package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmallrefundmessagesubmit 提交退款单留言
// taobao.openmall.refund.message.submit
//
// OpenMall业务提交退款单留言
func Taobaoopenmallrefundmessagesubmit(clt *core.SDKClient, req *openmall.TaobaoopenmallrefundmessagesubmitAPIRequest, session string) (*openmall.TaobaoopenmallrefundmessagesubmitAPIResponse, error) {
	var resp openmall.TaobaoopenmallrefundmessagesubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
