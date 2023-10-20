package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameinteractivegamestatusgetAPIResponse 获取游戏状态 API返回值
// alibaba.cloudgame.interactive.game.status.get
//
// 获取游戏状态
type AlibabacloudgameinteractivegamestatusgetAPIResponse struct {
	model.CommonResponse
	AlibabacloudgameinteractivegamestatusgetAPIResponseModel
}

// AlibabacloudgameinteractivegamestatusgetAPIResponseModel is 获取游戏状态 成功返回结果
type AlibabacloudgameinteractivegamestatusgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_interactive_game_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabacloudgameinteractivegamestatusgetResult `json:"result,omitempty" xml:"result,omitempty"`
}
