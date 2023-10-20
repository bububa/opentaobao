package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractWirelessDrawAPIResponse 双11到店互动无线端抽奖接口鉴权 API返回值
// alibaba.interact.wireless.draw
//
// 双11到店互动无线端mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 坯子
type AlibabaInteractWirelessDrawAPIResponse struct {
	model.CommonResponse
	AlibabaInteractWirelessDrawAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractWirelessDrawAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractWirelessDrawAPIResponseModel).Reset()
}

// AlibabaInteractWirelessDrawAPIResponseModel is 双11到店互动无线端抽奖接口鉴权 成功返回结果
type AlibabaInteractWirelessDrawAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_wireless_draw_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractWirelessDrawAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractWirelessDrawAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractWirelessDrawAPIResponse)
	},
}

// GetAlibabaInteractWirelessDrawAPIResponse 从 sync.Pool 获取 AlibabaInteractWirelessDrawAPIResponse
func GetAlibabaInteractWirelessDrawAPIResponse() *AlibabaInteractWirelessDrawAPIResponse {
	return poolAlibabaInteractWirelessDrawAPIResponse.Get().(*AlibabaInteractWirelessDrawAPIResponse)
}

// ReleaseAlibabaInteractWirelessDrawAPIResponse 将 AlibabaInteractWirelessDrawAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractWirelessDrawAPIResponse(v *AlibabaInteractWirelessDrawAPIResponse) {
	v.Reset()
	poolAlibabaInteractWirelessDrawAPIResponse.Put(v)
}
