package aliqin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcVoiceNumSinglecall 语音通知
// alibaba.aliqin.fc.voice.num.singlecall
//
// 向指定手机号码发起单向呼叫，播放指定的语音文件内容。使用前需要在阿里大于管理中心添加去电显示号码与语音文件。
func AlibabaAliqinFcVoiceNumSinglecall(ctx context.Context, clt *core.SDKClient, req *aliqin.AlibabaAliqinFcVoiceNumSinglecallAPIRequest, resp *aliqin.AlibabaAliqinFcVoiceNumSinglecallAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
