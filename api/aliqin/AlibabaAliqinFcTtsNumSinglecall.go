package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// Alibabaaliqinfcttsnumsinglecall 文本转语音通知
// alibaba.aliqin.fc.tts.num.singlecall
//
// 向指定手机号码发起单向呼叫，将文本模板内容转化为语音播放给被叫方。使用前需要在阿里大于管理中心添加去电显示号码与文本转语音模板。
func Alibabaaliqinfcttsnumsinglecall(clt *core.SDKClient, req *aliqin.AlibabaaliqinfcttsnumsinglecallAPIRequest, session string) (*aliqin.AlibabaaliqinfcttsnumsinglecallAPIResponse, error) {
	var resp aliqin.AlibabaaliqinfcttsnumsinglecallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
