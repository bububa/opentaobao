package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse 催促小B发送短信 API返回值
// alibaba.alihouse.newhome.adviser.message.notice
//
// 催促小B发送短信
type AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponseModel is 催促小B发送短信 成功返回结果
type AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_adviser_message_notice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeAdviserMessageNoticeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse
func GetAlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse() *AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse {
	return poolAlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse.Get().(*AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse 将 AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse(v *AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse.Put(v)
}
