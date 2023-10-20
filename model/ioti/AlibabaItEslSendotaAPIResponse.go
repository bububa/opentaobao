package ioti

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItEslSendotaAPIResponse 电子价签ota接口 API返回值
// alibaba.it.esl.sendota
//
// 厂测接口，电子价签ota接口
type AlibabaItEslSendotaAPIResponse struct {
	model.CommonResponse
	AlibabaItEslSendotaAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaItEslSendotaAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaItEslSendotaAPIResponseModel).Reset()
}

// AlibabaItEslSendotaAPIResponseModel is 电子价签ota接口 成功返回结果
type AlibabaItEslSendotaAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_it_esl_sendota_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Can not find Corresponding AP MAC with ESL
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaItEslSendotaAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaItEslSendotaAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaItEslSendotaAPIResponse)
	},
}

// GetAlibabaItEslSendotaAPIResponse 从 sync.Pool 获取 AlibabaItEslSendotaAPIResponse
func GetAlibabaItEslSendotaAPIResponse() *AlibabaItEslSendotaAPIResponse {
	return poolAlibabaItEslSendotaAPIResponse.Get().(*AlibabaItEslSendotaAPIResponse)
}

// ReleaseAlibabaItEslSendotaAPIResponse 将 AlibabaItEslSendotaAPIResponse 保存到 sync.Pool
func ReleaseAlibabaItEslSendotaAPIResponse(v *AlibabaItEslSendotaAPIResponse) {
	v.Reset()
	poolAlibabaItEslSendotaAPIResponse.Put(v)
}
