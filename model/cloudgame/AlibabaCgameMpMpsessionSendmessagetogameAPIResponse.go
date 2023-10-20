package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCgameMpMpsessionSendmessagetogameAPIResponse 发送消息给游戏 API返回值
// alibaba.cgame.mp.mpsession.sendmessagetogame
//
// 发送消息给游戏
type AlibabaCgameMpMpsessionSendmessagetogameAPIResponse struct {
	model.CommonResponse
	AlibabaCgameMpMpsessionSendmessagetogameAPIResponseModel
}

// AlibabaCgameMpMpsessionSendmessagetogameAPIResponseModel is 发送消息给游戏 成功返回结果
type AlibabaCgameMpMpsessionSendmessagetogameAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cgame_mp_mpsession_sendmessagetogame_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaCgameMpMpsessionSendmessagetogameResult `json:"result,omitempty" xml:"result,omitempty"`
}
