package ioti

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItEslEslimageShowimagecommonAPIResponse 对混合云提供的刷图接口 API返回值
// alibaba.it.esl.eslimage.showimagecommon
//
// 混合云使用，提供给isv和我们混合云环境部署的应用刷图
type AlibabaItEslEslimageShowimagecommonAPIResponse struct {
	model.CommonResponse
	AlibabaItEslEslimageShowimagecommonAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaItEslEslimageShowimagecommonAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaItEslEslimageShowimagecommonAPIResponseModel).Reset()
}

// AlibabaItEslEslimageShowimagecommonAPIResponseModel is 对混合云提供的刷图接口 成功返回结果
type AlibabaItEslEslimageShowimagecommonAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_it_esl_eslimage_showimagecommon_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultmsg
	Resultmsg string `json:"resultmsg,omitempty" xml:"resultmsg,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaItEslEslimageShowimagecommonAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Resultmsg = ""
}

var poolAlibabaItEslEslimageShowimagecommonAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaItEslEslimageShowimagecommonAPIResponse)
	},
}

// GetAlibabaItEslEslimageShowimagecommonAPIResponse 从 sync.Pool 获取 AlibabaItEslEslimageShowimagecommonAPIResponse
func GetAlibabaItEslEslimageShowimagecommonAPIResponse() *AlibabaItEslEslimageShowimagecommonAPIResponse {
	return poolAlibabaItEslEslimageShowimagecommonAPIResponse.Get().(*AlibabaItEslEslimageShowimagecommonAPIResponse)
}

// ReleaseAlibabaItEslEslimageShowimagecommonAPIResponse 将 AlibabaItEslEslimageShowimagecommonAPIResponse 保存到 sync.Pool
func ReleaseAlibabaItEslEslimageShowimagecommonAPIResponse(v *AlibabaItEslEslimageShowimagecommonAPIResponse) {
	v.Reset()
	poolAlibabaItEslEslimageShowimagecommonAPIResponse.Put(v)
}
