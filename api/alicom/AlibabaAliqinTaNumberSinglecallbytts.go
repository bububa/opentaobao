package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinTaNumberSinglecallbytts 根据号码tts单呼
// alibaba.aliqin.ta.number.singlecallbytts
//
// 将语音验证码和语音通知发布至聚石塔渠道
func AlibabaAliqinTaNumberSinglecallbytts(clt *core.SDKClient, req *alicom.AlibabaAliqinTaNumberSinglecallbyttsAPIRequest, resp *alicom.AlibabaAliqinTaNumberSinglecallbyttsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
