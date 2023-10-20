package cloudgame

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGameHeartbeatAPIResponse 游戏玩家心跳 API返回值
// alibaba.cloudgame.interactive.game.heartbeat
//
// 游戏玩家心跳
type AlibabaCloudgameInteractiveGameHeartbeatAPIResponse struct {
	model.CommonResponse
	AlibabaCloudgameInteractiveGameHeartbeatAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCloudgameInteractiveGameHeartbeatAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCloudgameInteractiveGameHeartbeatAPIResponseModel).Reset()
}

// AlibabaCloudgameInteractiveGameHeartbeatAPIResponseModel is 游戏玩家心跳 成功返回结果
type AlibabaCloudgameInteractiveGameHeartbeatAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_interactive_game_heartbeat_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaCloudgameInteractiveGameHeartbeatResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCloudgameInteractiveGameHeartbeatAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCloudgameInteractiveGameHeartbeatAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGameHeartbeatAPIResponse)
	},
}

// GetAlibabaCloudgameInteractiveGameHeartbeatAPIResponse 从 sync.Pool 获取 AlibabaCloudgameInteractiveGameHeartbeatAPIResponse
func GetAlibabaCloudgameInteractiveGameHeartbeatAPIResponse() *AlibabaCloudgameInteractiveGameHeartbeatAPIResponse {
	return poolAlibabaCloudgameInteractiveGameHeartbeatAPIResponse.Get().(*AlibabaCloudgameInteractiveGameHeartbeatAPIResponse)
}

// ReleaseAlibabaCloudgameInteractiveGameHeartbeatAPIResponse 将 AlibabaCloudgameInteractiveGameHeartbeatAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCloudgameInteractiveGameHeartbeatAPIResponse(v *AlibabaCloudgameInteractiveGameHeartbeatAPIResponse) {
	v.Reset()
	poolAlibabaCloudgameInteractiveGameHeartbeatAPIResponse.Put(v)
}
