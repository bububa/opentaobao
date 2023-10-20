package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameinteractivegamestartAPIResponse 云游戏场景互动开始游戏 API返回值
// alibaba.cloudgame.interactive.game.start
//
// 开始游戏
type AlibabacloudgameinteractivegamestartAPIResponse struct {
	model.CommonResponse
	AlibabacloudgameinteractivegamestartAPIResponseModel
}

// AlibabacloudgameinteractivegamestartAPIResponseModel is 云游戏场景互动开始游戏 成功返回结果
type AlibabacloudgameinteractivegamestartAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_interactive_game_start_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabacloudgameinteractivegamestartResult `json:"result,omitempty" xml:"result,omitempty"`
}
