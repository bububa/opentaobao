package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsMessageSend 聚石塔数据paas短信发送接口
// taobao.jst.sms.message.send
//
// 聚石塔短信PAAS场景中，ISV通过该API帮商家发送短信给用户。
func TaobaoJstSmsMessageSend(clt *core.SDKClient, req *jst.TaobaoJstSmsMessageSendAPIRequest, resp *jst.TaobaoJstSmsMessageSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
