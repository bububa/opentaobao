package cloudgame

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse 获取用户状态 API返回值
// alibaba.cloudgame.interactive.game.player.status.get
//
// 获取用户状态
type AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse struct {
	model.CommonResponse
	AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponseModel).Reset()
}

// AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponseModel is 获取用户状态 成功返回结果
type AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_interactive_game_player_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaCloudgameInteractiveGamePlayerStatusGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse)
	},
}

// GetAlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse 从 sync.Pool 获取 AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse
func GetAlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse() *AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse {
	return poolAlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse.Get().(*AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse)
}

// ReleaseAlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse 将 AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse(v *AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse) {
	v.Reset()
	poolAlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse.Put(v)
}
