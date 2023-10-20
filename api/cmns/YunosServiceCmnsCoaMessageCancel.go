package cmns

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// YunosServiceCmnsCoaMessageCancel CMNS消息撤回
// yunos.service.cmns.coa.message.cancel
//
// 此接口用户撤回之前已经发出去的消息，根据消息ID撤回，只能撤回此appKey创建的消息。
func YunosServiceCmnsCoaMessageCancel(clt *core.SDKClient, req *cmns.YunosServiceCmnsCoaMessageCancelAPIRequest, session string) (*cmns.YunosServiceCmnsCoaMessageCancelAPIResponse, error) {
	var resp cmns.YunosServiceCmnsCoaMessageCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
