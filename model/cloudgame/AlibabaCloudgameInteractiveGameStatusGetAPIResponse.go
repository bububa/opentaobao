package cloudgame

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGameStatusGetAPIResponse 获取游戏状态 API返回值
// alibaba.cloudgame.interactive.game.status.get
//
// 获取游戏状态
type AlibabaCloudgameInteractiveGameStatusGetAPIResponse struct {
	model.CommonResponse
	AlibabaCloudgameInteractiveGameStatusGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCloudgameInteractiveGameStatusGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCloudgameInteractiveGameStatusGetAPIResponseModel).Reset()
}

// AlibabaCloudgameInteractiveGameStatusGetAPIResponseModel is 获取游戏状态 成功返回结果
type AlibabaCloudgameInteractiveGameStatusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_interactive_game_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaCloudgameInteractiveGameStatusGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCloudgameInteractiveGameStatusGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCloudgameInteractiveGameStatusGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGameStatusGetAPIResponse)
	},
}

// GetAlibabaCloudgameInteractiveGameStatusGetAPIResponse 从 sync.Pool 获取 AlibabaCloudgameInteractiveGameStatusGetAPIResponse
func GetAlibabaCloudgameInteractiveGameStatusGetAPIResponse() *AlibabaCloudgameInteractiveGameStatusGetAPIResponse {
	return poolAlibabaCloudgameInteractiveGameStatusGetAPIResponse.Get().(*AlibabaCloudgameInteractiveGameStatusGetAPIResponse)
}

// ReleaseAlibabaCloudgameInteractiveGameStatusGetAPIResponse 将 AlibabaCloudgameInteractiveGameStatusGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCloudgameInteractiveGameStatusGetAPIResponse(v *AlibabaCloudgameInteractiveGameStatusGetAPIResponse) {
	v.Reset()
	poolAlibabaCloudgameInteractiveGameStatusGetAPIResponse.Put(v)
}
