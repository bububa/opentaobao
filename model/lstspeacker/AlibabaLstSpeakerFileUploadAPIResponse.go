package lstspeacker

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstSpeakerFileUploadAPIResponse 如意音箱音频文件长传 API返回值
// alibaba.lst.speaker.file.upload
//
// 如意音箱音频文件长传
type AlibabaLstSpeakerFileUploadAPIResponse struct {
	model.CommonResponse
	AlibabaLstSpeakerFileUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstSpeakerFileUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstSpeakerFileUploadAPIResponseModel).Reset()
}

// AlibabaLstSpeakerFileUploadAPIResponseModel is 如意音箱音频文件长传 成功返回结果
type AlibabaLstSpeakerFileUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_speaker_file_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabaLstSpeakerFileUploadResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstSpeakerFileUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstSpeakerFileUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstSpeakerFileUploadAPIResponse)
	},
}

// GetAlibabaLstSpeakerFileUploadAPIResponse 从 sync.Pool 获取 AlibabaLstSpeakerFileUploadAPIResponse
func GetAlibabaLstSpeakerFileUploadAPIResponse() *AlibabaLstSpeakerFileUploadAPIResponse {
	return poolAlibabaLstSpeakerFileUploadAPIResponse.Get().(*AlibabaLstSpeakerFileUploadAPIResponse)
}

// ReleaseAlibabaLstSpeakerFileUploadAPIResponse 将 AlibabaLstSpeakerFileUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstSpeakerFileUploadAPIResponse(v *AlibabaLstSpeakerFileUploadAPIResponse) {
	v.Reset()
	poolAlibabaLstSpeakerFileUploadAPIResponse.Put(v)
}
