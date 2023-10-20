package aliqin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotCardStatusAPIResponse 物联卡状态查询 API返回值
// alibaba.aliqin.fc.iot.cardStatus
//
// 物联卡状态查询
type AlibabaAliqinFcIotCardStatusAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIotCardStatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotCardStatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFcIotCardStatusAPIResponseModel).Reset()
}

// AlibabaAliqinFcIotCardStatusAPIResponseModel is 物联卡状态查询 成功返回结果
type AlibabaAliqinFcIotCardStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_cardStatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *AlibabaAliqinFcIotCardStatusResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotCardStatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinFcIotCardStatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotCardStatusAPIResponse)
	},
}

// GetAlibabaAliqinFcIotCardStatusAPIResponse 从 sync.Pool 获取 AlibabaAliqinFcIotCardStatusAPIResponse
func GetAlibabaAliqinFcIotCardStatusAPIResponse() *AlibabaAliqinFcIotCardStatusAPIResponse {
	return poolAlibabaAliqinFcIotCardStatusAPIResponse.Get().(*AlibabaAliqinFcIotCardStatusAPIResponse)
}

// ReleaseAlibabaAliqinFcIotCardStatusAPIResponse 将 AlibabaAliqinFcIotCardStatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFcIotCardStatusAPIResponse(v *AlibabaAliqinFcIotCardStatusAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFcIotCardStatusAPIResponse.Put(v)
}
