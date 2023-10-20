package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameinteractivegamejoincodeassignAPIResponse 分配joinCode API返回值
// alibaba.cloudgame.interactive.game.joincode.assign
//
// 分配joinCode
type AlibabacloudgameinteractivegamejoincodeassignAPIResponse struct {
	model.CommonResponse
	AlibabacloudgameinteractivegamejoincodeassignAPIResponseModel
}

// AlibabacloudgameinteractivegamejoincodeassignAPIResponseModel is 分配joinCode 成功返回结果
type AlibabacloudgameinteractivegamejoincodeassignAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_interactive_game_joincode_assign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabacloudgameinteractivegamejoincodeassignResult `json:"result,omitempty" xml:"result,omitempty"`
}
