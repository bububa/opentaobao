package wdkitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPictureUploadAPIResponse 图片上传接口 API返回值
// alibaba.wdk.picture.upload
//
// 上传图片
type AlibabaWdkPictureUploadAPIResponse struct {
	model.CommonResponse
	AlibabaWdkPictureUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkPictureUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkPictureUploadAPIResponseModel).Reset()
}

// AlibabaWdkPictureUploadAPIResponseModel is 图片上传接口 成功返回结果
type AlibabaWdkPictureUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_picture_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// apiresult
	Result *AlibabaWdkPictureUploadApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkPictureUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkPictureUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkPictureUploadAPIResponse)
	},
}

// GetAlibabaWdkPictureUploadAPIResponse 从 sync.Pool 获取 AlibabaWdkPictureUploadAPIResponse
func GetAlibabaWdkPictureUploadAPIResponse() *AlibabaWdkPictureUploadAPIResponse {
	return poolAlibabaWdkPictureUploadAPIResponse.Get().(*AlibabaWdkPictureUploadAPIResponse)
}

// ReleaseAlibabaWdkPictureUploadAPIResponse 将 AlibabaWdkPictureUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkPictureUploadAPIResponse(v *AlibabaWdkPictureUploadAPIResponse) {
	v.Reset()
	poolAlibabaWdkPictureUploadAPIResponse.Put(v)
}
