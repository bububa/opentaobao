package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGameStartAPIResponse 云游戏场景互动开始游戏 API返回值
// alibaba.cloudgame.interactive.game.start
//
// 开始游戏
type AlibabaCloudgameInteractiveGameStartAPIResponse struct {
	model.CommonResponse
	AlibabaCloudgameInteractiveGameStartAPIResponseModel
}

// AlibabaCloudgameInteractiveGameStartAPIResponseModel is 云游戏场景互动开始游戏 成功返回结果
type AlibabaCloudgameInteractiveGameStartAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_interactive_game_start_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaCloudgameInteractiveGameStartResult `json:"result,omitempty" xml:"result,omitempty"`
}
