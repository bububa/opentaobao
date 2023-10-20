package cloudgame

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCgameAvatarUserbodyQueryAPIResponse 用户Avatar body查询 API返回值
// alibaba.cgame.avatar.userbody.query
//
// Avatar用户body数据查询
type AlibabaCgameAvatarUserbodyQueryAPIResponse struct {
	model.CommonResponse
	AlibabaCgameAvatarUserbodyQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCgameAvatarUserbodyQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCgameAvatarUserbodyQueryAPIResponseModel).Reset()
}

// AlibabaCgameAvatarUserbodyQueryAPIResponseModel is 用户Avatar body查询 成功返回结果
type AlibabaCgameAvatarUserbodyQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cgame_avatar_userbody_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaCgameAvatarUserbodyQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCgameAvatarUserbodyQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCgameAvatarUserbodyQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCgameAvatarUserbodyQueryAPIResponse)
	},
}

// GetAlibabaCgameAvatarUserbodyQueryAPIResponse 从 sync.Pool 获取 AlibabaCgameAvatarUserbodyQueryAPIResponse
func GetAlibabaCgameAvatarUserbodyQueryAPIResponse() *AlibabaCgameAvatarUserbodyQueryAPIResponse {
	return poolAlibabaCgameAvatarUserbodyQueryAPIResponse.Get().(*AlibabaCgameAvatarUserbodyQueryAPIResponse)
}

// ReleaseAlibabaCgameAvatarUserbodyQueryAPIResponse 将 AlibabaCgameAvatarUserbodyQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCgameAvatarUserbodyQueryAPIResponse(v *AlibabaCgameAvatarUserbodyQueryAPIResponse) {
	v.Reset()
	poolAlibabaCgameAvatarUserbodyQueryAPIResponse.Put(v)
}
