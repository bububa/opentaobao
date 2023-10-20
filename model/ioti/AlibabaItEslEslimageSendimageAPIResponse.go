package ioti

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItEslEslimageSendimageAPIResponse 下发厂测初始化图片 API返回值
// alibaba.it.esl.eslimage.sendimage
//
// 工厂对生产出的电子价签进行全流程功能测试，能将出场图片通过ESL系统初始化到电子价签中
type AlibabaItEslEslimageSendimageAPIResponse struct {
	model.CommonResponse
	AlibabaItEslEslimageSendimageAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaItEslEslimageSendimageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaItEslEslimageSendimageAPIResponseModel).Reset()
}

// AlibabaItEslEslimageSendimageAPIResponseModel is 下发厂测初始化图片 成功返回结果
type AlibabaItEslEslimageSendimageAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_it_esl_eslimage_sendimage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultmsg
	Resultmsg string `json:"resultmsg,omitempty" xml:"resultmsg,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaItEslEslimageSendimageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Resultmsg = ""
}

var poolAlibabaItEslEslimageSendimageAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaItEslEslimageSendimageAPIResponse)
	},
}

// GetAlibabaItEslEslimageSendimageAPIResponse 从 sync.Pool 获取 AlibabaItEslEslimageSendimageAPIResponse
func GetAlibabaItEslEslimageSendimageAPIResponse() *AlibabaItEslEslimageSendimageAPIResponse {
	return poolAlibabaItEslEslimageSendimageAPIResponse.Get().(*AlibabaItEslEslimageSendimageAPIResponse)
}

// ReleaseAlibabaItEslEslimageSendimageAPIResponse 将 AlibabaItEslEslimageSendimageAPIResponse 保存到 sync.Pool
func ReleaseAlibabaItEslEslimageSendimageAPIResponse(v *AlibabaItEslEslimageSendimageAPIResponse) {
	v.Reset()
	poolAlibabaItEslEslimageSendimageAPIResponse.Put(v)
}
