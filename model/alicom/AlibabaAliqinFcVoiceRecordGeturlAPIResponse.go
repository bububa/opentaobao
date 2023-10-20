package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinfcvoicerecordgeturlAPIResponse 录音文件下载 API返回值
// alibaba.aliqin.fc.voice.record.geturl
//
// 完成录音文件的下载地址获取操作
type AlibabaaliqinfcvoicerecordgeturlAPIResponse struct {
	model.CommonResponse
	AlibabaaliqinfcvoicerecordgeturlAPIResponseModel
}

// AlibabaaliqinfcvoicerecordgeturlAPIResponseModel is 录音文件下载 成功返回结果
type AlibabaaliqinfcvoicerecordgeturlAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_voice_record_geturl_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaaliqinfcvoicerecordgeturlResult `json:"result,omitempty" xml:"result,omitempty"`
}
