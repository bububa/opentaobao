package aliqin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotCardofferAPIResponse 查询物联网卡上订购的offer API返回值
// alibaba.aliqin.fc.iot.cardoffer
//
// 查询物联网卡上订购的offer
type AlibabaAliqinFcIotCardofferAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIotCardofferAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotCardofferAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFcIotCardofferAPIResponseModel).Reset()
}

// AlibabaAliqinFcIotCardofferAPIResponseModel is 查询物联网卡上订购的offer 成功返回结果
type AlibabaAliqinFcIotCardofferAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_cardoffer_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *AlibabaAliqinFcIotCardofferResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotCardofferAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinFcIotCardofferAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotCardofferAPIResponse)
	},
}

// GetAlibabaAliqinFcIotCardofferAPIResponse 从 sync.Pool 获取 AlibabaAliqinFcIotCardofferAPIResponse
func GetAlibabaAliqinFcIotCardofferAPIResponse() *AlibabaAliqinFcIotCardofferAPIResponse {
	return poolAlibabaAliqinFcIotCardofferAPIResponse.Get().(*AlibabaAliqinFcIotCardofferAPIResponse)
}

// ReleaseAlibabaAliqinFcIotCardofferAPIResponse 将 AlibabaAliqinFcIotCardofferAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFcIotCardofferAPIResponse(v *AlibabaAliqinFcIotCardofferAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFcIotCardofferAPIResponse.Put(v)
}
