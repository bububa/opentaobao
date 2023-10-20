package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacgameavataruserbodyqueryAPIResponse 用户Avatar body查询 API返回值
// alibaba.cgame.avatar.userbody.query
//
// Avatar用户body数据查询
type AlibabacgameavataruserbodyqueryAPIResponse struct {
	model.CommonResponse
	AlibabacgameavataruserbodyqueryAPIResponseModel
}

// AlibabacgameavataruserbodyqueryAPIResponseModel is 用户Avatar body查询 成功返回结果
type AlibabacgameavataruserbodyqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cgame_avatar_userbody_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabacgameavataruserbodyqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
