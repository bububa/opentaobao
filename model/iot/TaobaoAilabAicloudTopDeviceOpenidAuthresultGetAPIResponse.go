package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse 获取openId设备授权码验证结果 API返回值
// taobao.ailab.aicloud.top.device.openid.authresult.get
//
// 获取openId设备授权码验证结果
type TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponseModel is 获取openId设备授权码验证结果 成功返回结果
type TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_openid_authresult_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse
func GetTaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse() *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse 将 TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse(v *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse.Put(v)
}
