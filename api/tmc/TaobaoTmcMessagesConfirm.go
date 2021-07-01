package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

/* TaobaoTmcMessagesConfirm
确认消费消息的状态
taobao.tmc.messages.confirm

确认消费消息的状态 */
func TaobaoTmcMessagesConfirm(clt *core.SDKClient, req *tmc.TaobaoTmcMessagesConfirmAPIRequest, session string) (*tmc.TaobaoTmcMessagesConfirmAPIResponse, error) {
	var resp tmc.TaobaoTmcMessagesConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
