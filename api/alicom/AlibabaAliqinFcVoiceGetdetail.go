package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

/* AlibabaAliqinFcVoiceGetdetail
获取呼叫详情
alibaba.aliqin.fc.voice.getdetail

通过呼叫id获取呼叫相关的数据 */
func AlibabaAliqinFcVoiceGetdetail(clt *core.SDKClient, req *alicom.AlibabaAliqinFcVoiceGetdetailAPIRequest, session string) (*alicom.AlibabaAliqinFcVoiceGetdetailAPIResponse, error) {
	var resp alicom.AlibabaAliqinFcVoiceGetdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
