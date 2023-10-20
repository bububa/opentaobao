package cloudgame

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGameRoomCreateAPIResponse 建游戏房间 API返回值
// alibaba.cloudgame.interactive.game.room.create
//
// 建游戏房间
type AlibabaCloudgameInteractiveGameRoomCreateAPIResponse struct {
	model.CommonResponse
	AlibabaCloudgameInteractiveGameRoomCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCloudgameInteractiveGameRoomCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCloudgameInteractiveGameRoomCreateAPIResponseModel).Reset()
}

// AlibabaCloudgameInteractiveGameRoomCreateAPIResponseModel is 建游戏房间 成功返回结果
type AlibabaCloudgameInteractiveGameRoomCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_interactive_game_room_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaCloudgameInteractiveGameRoomCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCloudgameInteractiveGameRoomCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCloudgameInteractiveGameRoomCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGameRoomCreateAPIResponse)
	},
}

// GetAlibabaCloudgameInteractiveGameRoomCreateAPIResponse 从 sync.Pool 获取 AlibabaCloudgameInteractiveGameRoomCreateAPIResponse
func GetAlibabaCloudgameInteractiveGameRoomCreateAPIResponse() *AlibabaCloudgameInteractiveGameRoomCreateAPIResponse {
	return poolAlibabaCloudgameInteractiveGameRoomCreateAPIResponse.Get().(*AlibabaCloudgameInteractiveGameRoomCreateAPIResponse)
}

// ReleaseAlibabaCloudgameInteractiveGameRoomCreateAPIResponse 将 AlibabaCloudgameInteractiveGameRoomCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCloudgameInteractiveGameRoomCreateAPIResponse(v *AlibabaCloudgameInteractiveGameRoomCreateAPIResponse) {
	v.Reset()
	poolAlibabaCloudgameInteractiveGameRoomCreateAPIResponse.Put(v)
}
