package shenjing

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIbShenjingVisitorPadUploadfaceAPIResponse 访客PAD上传人脸 API返回值
// alibaba.ib.shenjing.visitor.pad.uploadface
//
// 访客PAD端上传人脸。
type AlibabaIbShenjingVisitorPadUploadfaceAPIResponse struct {
	model.CommonResponse
	AlibabaIbShenjingVisitorPadUploadfaceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIbShenjingVisitorPadUploadfaceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIbShenjingVisitorPadUploadfaceAPIResponseModel).Reset()
}

// AlibabaIbShenjingVisitorPadUploadfaceAPIResponseModel is 访客PAD上传人脸 成功返回结果
type AlibabaIbShenjingVisitorPadUploadfaceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ib_shenjing_visitor_pad_uploadface_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// request_id
	ResultRequestId string `json:"result_request_id,omitempty" xml:"result_request_id,omitempty"`
	// 内容
	Content *UploadFaceDo `json:"content,omitempty" xml:"content,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIbShenjingVisitorPadUploadfaceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.ResultRequestId = ""
	m.Content = nil
}

var poolAlibabaIbShenjingVisitorPadUploadfaceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIbShenjingVisitorPadUploadfaceAPIResponse)
	},
}

// GetAlibabaIbShenjingVisitorPadUploadfaceAPIResponse 从 sync.Pool 获取 AlibabaIbShenjingVisitorPadUploadfaceAPIResponse
func GetAlibabaIbShenjingVisitorPadUploadfaceAPIResponse() *AlibabaIbShenjingVisitorPadUploadfaceAPIResponse {
	return poolAlibabaIbShenjingVisitorPadUploadfaceAPIResponse.Get().(*AlibabaIbShenjingVisitorPadUploadfaceAPIResponse)
}

// ReleaseAlibabaIbShenjingVisitorPadUploadfaceAPIResponse 将 AlibabaIbShenjingVisitorPadUploadfaceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIbShenjingVisitorPadUploadfaceAPIResponse(v *AlibabaIbShenjingVisitorPadUploadfaceAPIResponse) {
	v.Reset()
	poolAlibabaIbShenjingVisitorPadUploadfaceAPIResponse.Put(v)
}
