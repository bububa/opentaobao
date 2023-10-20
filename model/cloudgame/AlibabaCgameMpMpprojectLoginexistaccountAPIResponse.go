package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacgamempmpprojectloginexistaccountAPIResponse 登录存在账号 API返回值
// alibaba.cgame.mp.mpproject.loginexistaccount
//
// 发送消息给游戏
type AlibabacgamempmpprojectloginexistaccountAPIResponse struct {
	model.CommonResponse
	AlibabacgamempmpprojectloginexistaccountAPIResponseModel
}

// AlibabacgamempmpprojectloginexistaccountAPIResponseModel is 登录存在账号 成功返回结果
type AlibabacgamempmpprojectloginexistaccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cgame_mp_mpproject_loginexistaccount_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabacgamempmpprojectloginexistaccountResult `json:"result,omitempty" xml:"result,omitempty"`
}
