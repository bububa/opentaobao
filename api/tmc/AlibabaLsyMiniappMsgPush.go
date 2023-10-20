package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// AlibabaLsyMiniappMsgPush 零售云小程序消息推送
// alibaba.lsy.miniapp.msg.push
//
// 零售云小程序消息推送，推送消息至零售云（喵零等）
func AlibabaLsyMiniappMsgPush(clt *core.SDKClient, req *tmc.AlibabaLsyMiniappMsgPushAPIRequest, resp *tmc.AlibabaLsyMiniappMsgPushAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
