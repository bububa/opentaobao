package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIotDeviceCorpusGetAPIResponse IoT设备支持语料获取 API返回值
// alibaba.iot.device.corpus.get
//
// ISV通过该接口获取天猫精灵IoT设备支持控制或查询的语料
type AlibabaIotDeviceCorpusGetAPIResponse struct {
	model.CommonResponse
	AlibabaIotDeviceCorpusGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIotDeviceCorpusGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIotDeviceCorpusGetAPIResponseModel).Reset()
}

// AlibabaIotDeviceCorpusGetAPIResponseModel is IoT设备支持语料获取 成功返回结果
type AlibabaIotDeviceCorpusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_iot_device_corpus_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	RetValues []DeviceCorpusTopDto `json:"ret_values,omitempty" xml:"ret_values>device_corpus_top_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIotDeviceCorpusGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetValues = m.RetValues[:0]
}

var poolAlibabaIotDeviceCorpusGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIotDeviceCorpusGetAPIResponse)
	},
}

// GetAlibabaIotDeviceCorpusGetAPIResponse 从 sync.Pool 获取 AlibabaIotDeviceCorpusGetAPIResponse
func GetAlibabaIotDeviceCorpusGetAPIResponse() *AlibabaIotDeviceCorpusGetAPIResponse {
	return poolAlibabaIotDeviceCorpusGetAPIResponse.Get().(*AlibabaIotDeviceCorpusGetAPIResponse)
}

// ReleaseAlibabaIotDeviceCorpusGetAPIResponse 将 AlibabaIotDeviceCorpusGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIotDeviceCorpusGetAPIResponse(v *AlibabaIotDeviceCorpusGetAPIResponse) {
	v.Reset()
	poolAlibabaIotDeviceCorpusGetAPIResponse.Put(v)
}
