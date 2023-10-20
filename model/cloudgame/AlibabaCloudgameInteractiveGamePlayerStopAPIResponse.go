package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameinteractivegameplayerstopAPIResponse 用户停止游戏 API返回值
// alibaba.cloudgame.interactive.game.player.stop
//
// 用户停止游戏
type AlibabacloudgameinteractivegameplayerstopAPIResponse struct {
	model.CommonResponse
	AlibabacloudgameinteractivegameplayerstopAPIResponseModel
}

// AlibabacloudgameinteractivegameplayerstopAPIResponseModel is 用户停止游戏 成功返回结果
type AlibabacloudgameinteractivegameplayerstopAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_interactive_game_player_stop_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabacloudgameinteractivegameplayerstopResult `json:"result,omitempty" xml:"result,omitempty"`
}
