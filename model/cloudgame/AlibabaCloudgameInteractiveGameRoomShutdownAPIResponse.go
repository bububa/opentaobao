package cloudgame

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGameRoomShutdownAPIResponse 强制关闭房间 API返回值
// alibaba.cloudgame.interactive.game.room.shutdown
//
// 强制关闭房间
type AlibabaCloudgameInteractiveGameRoomShutdownAPIResponse struct {
	model.CommonResponse
	AlibabaCloudgameInteractiveGameRoomShutdownAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCloudgameInteractiveGameRoomShutdownAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCloudgameInteractiveGameRoomShutdownAPIResponseModel).Reset()
}

// AlibabaCloudgameInteractiveGameRoomShutdownAPIResponseModel is 强制关闭房间 成功返回结果
type AlibabaCloudgameInteractiveGameRoomShutdownAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_interactive_game_room_shutdown_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaCloudgameInteractiveGameRoomShutdownResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCloudgameInteractiveGameRoomShutdownAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCloudgameInteractiveGameRoomShutdownAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGameRoomShutdownAPIResponse)
	},
}

// GetAlibabaCloudgameInteractiveGameRoomShutdownAPIResponse 从 sync.Pool 获取 AlibabaCloudgameInteractiveGameRoomShutdownAPIResponse
func GetAlibabaCloudgameInteractiveGameRoomShutdownAPIResponse() *AlibabaCloudgameInteractiveGameRoomShutdownAPIResponse {
	return poolAlibabaCloudgameInteractiveGameRoomShutdownAPIResponse.Get().(*AlibabaCloudgameInteractiveGameRoomShutdownAPIResponse)
}

// ReleaseAlibabaCloudgameInteractiveGameRoomShutdownAPIResponse 将 AlibabaCloudgameInteractiveGameRoomShutdownAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCloudgameInteractiveGameRoomShutdownAPIResponse(v *AlibabaCloudgameInteractiveGameRoomShutdownAPIResponse) {
	v.Reset()
	poolAlibabaCloudgameInteractiveGameRoomShutdownAPIResponse.Put(v)
}
