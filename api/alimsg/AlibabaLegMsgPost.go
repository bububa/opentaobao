package alimsg

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimsg"
)

/* AlibabaLegMsgPost
集团法务消息发送
alibaba.leg.msg.post

消息发送能力 */
func AlibabaLegMsgPost(clt *core.SDKClient, req *alimsg.AlibabaLegMsgPostAPIRequest, session string) (*alimsg.AlibabaLegMsgPostAPIResponse, error) {
	var resp alimsg.AlibabaLegMsgPostAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
