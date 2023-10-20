package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// Alibabaaliqinfcvoicenumsinglecall 语音通知
// alibaba.aliqin.fc.voice.num.singlecall
//
// 向指定手机号码发起单向呼叫，播放指定的语音文件内容。使用前需要在阿里大于管理中心添加去电显示号码与语音文件。
func Alibabaaliqinfcvoicenumsinglecall(clt *core.SDKClient, req *aliqin.AlibabaaliqinfcvoicenumsinglecallAPIRequest, session string) (*aliqin.AlibabaaliqinfcvoicenumsinglecallAPIResponse, error) {
	var resp aliqin.AlibabaaliqinfcvoicenumsinglecallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
