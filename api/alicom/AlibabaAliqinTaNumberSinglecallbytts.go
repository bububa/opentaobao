package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

/* AlibabaAliqinTaNumberSinglecallbytts
根据号码tts单呼
alibaba.aliqin.ta.number.singlecallbytts

将语音验证码和语音通知发布至聚石塔渠道 */
func AlibabaAliqinTaNumberSinglecallbytts(clt *core.SDKClient, req *alicom.AlibabaAliqinTaNumberSinglecallbyttsAPIRequest, session string) (*alicom.AlibabaAliqinTaNumberSinglecallbyttsAPIResponse, error) {
	var resp alicom.AlibabaAliqinTaNumberSinglecallbyttsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
