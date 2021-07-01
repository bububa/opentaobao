package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcVoiceRecordGeturlAPIRequest
录音文件下载 API请求
alibaba.aliqin.fc.voice.record.geturl

完成录音文件的下载地址获取操作 */
type AlibabaAliqinFcVoiceRecordGeturlAPIRequest struct {
	model.Params
	// 一次通话的唯一标识id
	_callId string
}

// New
