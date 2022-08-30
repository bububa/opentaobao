package cloudgame

import (
	"encoding/xml"

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

// AlibabaCloudgameInteractiveGameRoomShutdownAPIResponseModel is 强制关闭房间 成功返回结果
type AlibabaCloudgameInteractiveGameRoomShutdownAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_interactive_game_room_shutdown_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaCloudgameInteractiveGameRoomShutdownResult `json:"result,omitempty" xml:"result,omitempty"`
}
