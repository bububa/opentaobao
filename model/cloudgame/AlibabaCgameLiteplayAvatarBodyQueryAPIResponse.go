package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacgameliteplayavatarbodyqueryAPIResponse 新氢玩Avatar脸部装扮数据查询 API返回值
// alibaba.cgame.liteplay.avatar.body.query
//
// 云游戏, 新氢玩, 围观互动,自研游戏包, 需要查询Avatar虚拟形象捏脸和装扮数据, 用来初始化游戏包形象.
type AlibabacgameliteplayavatarbodyqueryAPIResponse struct {
	model.CommonResponse
	AlibabacgameliteplayavatarbodyqueryAPIResponseModel
}

// AlibabacgameliteplayavatarbodyqueryAPIResponseModel is 新氢玩Avatar脸部装扮数据查询 成功返回结果
type AlibabacgameliteplayavatarbodyqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cgame_liteplay_avatar_body_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabacgameliteplayavatarbodyqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
