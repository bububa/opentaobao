package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse 获取设备授权码验证结果 API返回值
// taobao.ailab.aicloud.top.device.authresult.get
//
// 获取设备授权码验证结果
type TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponseModel is 获取设备授权码验证结果 成功返回结果
type TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_authresult_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse
func GetTaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse() *TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse 将 TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse(v *TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse.Put(v)
}
