package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCgameAvatarUserbodyQueryAPIResponse
用户Avatar body查询 API返回值
alibaba.cgame.avatar.userbody.query

Avatar用户body数据查询 */
type AlibabaCgameAvatarUserbodyQueryAPIResponse struct {
	model.CommonResponse
	AlibabaCgameAvatarUserbodyQueryAPIResponseModel
}

// AlibabaCgameAvatarUserbodyQueryAPIResponseModel is 用户Avatar body查询 成功返回结果
type AlibabaCgameAvatarUserbodyQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cgame_avatar_userbody_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaCgameAvatarUserbodyQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
