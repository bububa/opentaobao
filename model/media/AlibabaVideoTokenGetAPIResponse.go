package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaVideoTokenGetAPIResponse 获取上传token API返回值
// alibaba.video.token.get
//
// 获取上传token
type AlibabaVideoTokenGetAPIResponse struct {
	model.CommonResponse
	AlibabaVideoTokenGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaVideoTokenGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaVideoTokenGetAPIResponseModel).Reset()
}

// AlibabaVideoTokenGetAPIResponseModel is 获取上传token 成功返回结果
type AlibabaVideoTokenGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_video_token_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 上传token
	UploadToken string `json:"upload_token,omitempty" xml:"upload_token,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaVideoTokenGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UploadToken = ""
}

var poolAlibabaVideoTokenGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaVideoTokenGetAPIResponse)
	},
}

// GetAlibabaVideoTokenGetAPIResponse 从 sync.Pool 获取 AlibabaVideoTokenGetAPIResponse
func GetAlibabaVideoTokenGetAPIResponse() *AlibabaVideoTokenGetAPIResponse {
	return poolAlibabaVideoTokenGetAPIResponse.Get().(*AlibabaVideoTokenGetAPIResponse)
}

// ReleaseAlibabaVideoTokenGetAPIResponse 将 AlibabaVideoTokenGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaVideoTokenGetAPIResponse(v *AlibabaVideoTokenGetAPIResponse) {
	v.Reset()
	poolAlibabaVideoTokenGetAPIResponse.Put(v)
}
