package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuPhotobankUploadAPIResponse 图片银行图片上传开放接口 API返回值
// alibaba.icbu.photobank.upload
//
// 图片银行图片上传开放接口
type AlibabaIcbuPhotobankUploadAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuPhotobankUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuPhotobankUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuPhotobankUploadAPIResponseModel).Reset()
}

// AlibabaIcbuPhotobankUploadAPIResponseModel is 图片银行图片上传开放接口 成功返回结果
type AlibabaIcbuPhotobankUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_photobank_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 图片信息
	UploadImageResponse *UploadImageResponseDo `json:"upload_image_response,omitempty" xml:"upload_image_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuPhotobankUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UploadImageResponse = nil
}

var poolAlibabaIcbuPhotobankUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuPhotobankUploadAPIResponse)
	},
}

// GetAlibabaIcbuPhotobankUploadAPIResponse 从 sync.Pool 获取 AlibabaIcbuPhotobankUploadAPIResponse
func GetAlibabaIcbuPhotobankUploadAPIResponse() *AlibabaIcbuPhotobankUploadAPIResponse {
	return poolAlibabaIcbuPhotobankUploadAPIResponse.Get().(*AlibabaIcbuPhotobankUploadAPIResponse)
}

// ReleaseAlibabaIcbuPhotobankUploadAPIResponse 将 AlibabaIcbuPhotobankUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuPhotobankUploadAPIResponse(v *AlibabaIcbuPhotobankUploadAPIResponse) {
	v.Reset()
	poolAlibabaIcbuPhotobankUploadAPIResponse.Put(v)
}
