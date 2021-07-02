package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimCustmsgPush 推送自定义openim消息
// taobao.openim.custmsg.push
//
// isv通过该接口给openim用户推送自定义消息
func TaobaoOpenimCustmsgPush(clt *core.SDKClient, req *openim.TaobaoOpenimCustmsgPushAPIRequest, session string) (*openim.TaobaoOpenimCustmsgPushAPIResponse, error) {
	var resp openim.TaobaoOpenimCustmsgPushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
