package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacgamempmpprojectinitnewprojectAPIResponse 创建新的mpproject API返回值
// alibaba.cgame.mp.mpproject.initnewproject
//
// 发送消息给游戏
type AlibabacgamempmpprojectinitnewprojectAPIResponse struct {
	model.CommonResponse
	AlibabacgamempmpprojectinitnewprojectAPIResponseModel
}

// AlibabacgamempmpprojectinitnewprojectAPIResponseModel is 创建新的mpproject 成功返回结果
type AlibabacgamempmpprojectinitnewprojectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cgame_mp_mpproject_initnewproject_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabacgamempmpprojectinitnewprojectResult `json:"result,omitempty" xml:"result,omitempty"`
}
