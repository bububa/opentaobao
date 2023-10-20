package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkRtaConsumerMatch 淘宝客-推广者-定向活动目标发布
// taobao.tbk.rta.consumer.match
//
// 淘客计划向用户推送某个定向活动时，调用该接口判断用户是否符合活动目标（淘客接入前需签署协议 https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
func TaobaoTbkRtaConsumerMatch(clt *core.SDKClient, req *tbk.TaobaoTbkRtaConsumerMatchAPIRequest, resp *tbk.TaobaoTbkRtaConsumerMatchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
