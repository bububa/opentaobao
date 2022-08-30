package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsOaidMessageSend 基于OAID的短信发送接口
// taobao.jst.sms.oaid.message.send
//
// 基于OAID的短信发送接口
func TaobaoJstSmsOaidMessageSend(clt *core.SDKClient, req *jst.TaobaoJstSmsOaidMessageSendAPIRequest, session string) (*jst.TaobaoJstSmsOaidMessageSendAPIResponse, error) {
	var resp jst.TaobaoJstSmsOaidMessageSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
