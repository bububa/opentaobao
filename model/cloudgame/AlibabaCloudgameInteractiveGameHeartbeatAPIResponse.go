package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameinteractivegameheartbeatAPIResponse 游戏玩家心跳 API返回值
// alibaba.cloudgame.interactive.game.heartbeat
//
// 游戏玩家心跳
type AlibabacloudgameinteractivegameheartbeatAPIResponse struct {
	model.CommonResponse
	AlibabacloudgameinteractivegameheartbeatAPIResponseModel
}

// AlibabacloudgameinteractivegameheartbeatAPIResponseModel is 游戏玩家心跳 成功返回结果
type AlibabacloudgameinteractivegameheartbeatAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_interactive_game_heartbeat_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabacloudgameinteractivegameheartbeatResult `json:"result,omitempty" xml:"result,omitempty"`
}
