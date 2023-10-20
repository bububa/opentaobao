package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinFcVoiceRecordGeturl 录音文件下载
// alibaba.aliqin.fc.voice.record.geturl
//
// 完成录音文件的下载地址获取操作
func AlibabaAliqinFcVoiceRecordGeturl(clt *core.SDKClient, req *alicom.AlibabaAliqinFcVoiceRecordGeturlAPIRequest, resp *alicom.AlibabaAliqinFcVoiceRecordGeturlAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
