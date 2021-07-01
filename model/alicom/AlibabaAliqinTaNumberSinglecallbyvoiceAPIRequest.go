package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest
根据号码tts单呼 API请求
alibaba.aliqin.ta.number.singlecallbyvoice

根据号码语音单呼 */
type AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest struct {
	model.Params
	// 单呼号码
	_calledNum string
	// 显示号码
	_calledShowNum string
	// 语音文件code
	_voiceCode string
	// 上下文参数 示例:{"extend":"回传参数"} extend为扩展信息作为回传参数的key
	_params string
}

// New
