package aliqin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotModbindAPIResponse 物联网绑定/换绑API API返回值
// alibaba.aliqin.fc.iot.modbind
//
// 支持用户的设备的换绑和解绑操作
type AlibabaAliqinFcIotModbindAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIotModbindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotModbindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFcIotModbindAPIResponseModel).Reset()
}

// AlibabaAliqinFcIotModbindAPIResponseModel is 物联网绑定/换绑API 成功返回结果
type AlibabaAliqinFcIotModbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_modbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinFcIotModbindResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotModbindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinFcIotModbindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotModbindAPIResponse)
	},
}

// GetAlibabaAliqinFcIotModbindAPIResponse 从 sync.Pool 获取 AlibabaAliqinFcIotModbindAPIResponse
func GetAlibabaAliqinFcIotModbindAPIResponse() *AlibabaAliqinFcIotModbindAPIResponse {
	return poolAlibabaAliqinFcIotModbindAPIResponse.Get().(*AlibabaAliqinFcIotModbindAPIResponse)
}

// ReleaseAlibabaAliqinFcIotModbindAPIResponse 将 AlibabaAliqinFcIotModbindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFcIotModbindAPIResponse(v *AlibabaAliqinFcIotModbindAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFcIotModbindAPIResponse.Put(v)
}
