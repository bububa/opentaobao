package cloudgame

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGameGamepadGetAPIResponse 获取虚拟手柄配置 API返回值
// alibaba.cloudgame.interactive.game.gamepad.get
//
// 获取虚拟手柄配置
type AlibabaCloudgameInteractiveGameGamepadGetAPIResponse struct {
	model.CommonResponse
	AlibabaCloudgameInteractiveGameGamepadGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCloudgameInteractiveGameGamepadGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCloudgameInteractiveGameGamepadGetAPIResponseModel).Reset()
}

// AlibabaCloudgameInteractiveGameGamepadGetAPIResponseModel is 获取虚拟手柄配置 成功返回结果
type AlibabaCloudgameInteractiveGameGamepadGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_interactive_game_gamepad_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *AlibabaCloudgameInteractiveGameGamepadGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCloudgameInteractiveGameGamepadGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCloudgameInteractiveGameGamepadGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGameGamepadGetAPIResponse)
	},
}

// GetAlibabaCloudgameInteractiveGameGamepadGetAPIResponse 从 sync.Pool 获取 AlibabaCloudgameInteractiveGameGamepadGetAPIResponse
func GetAlibabaCloudgameInteractiveGameGamepadGetAPIResponse() *AlibabaCloudgameInteractiveGameGamepadGetAPIResponse {
	return poolAlibabaCloudgameInteractiveGameGamepadGetAPIResponse.Get().(*AlibabaCloudgameInteractiveGameGamepadGetAPIResponse)
}

// ReleaseAlibabaCloudgameInteractiveGameGamepadGetAPIResponse 将 AlibabaCloudgameInteractiveGameGamepadGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCloudgameInteractiveGameGamepadGetAPIResponse(v *AlibabaCloudgameInteractiveGameGamepadGetAPIResponse) {
	v.Reset()
	poolAlibabaCloudgameInteractiveGameGamepadGetAPIResponse.Put(v)
}
