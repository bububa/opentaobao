package cloudgame

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCgameMpMpprojectLoginexistaccountAPIResponse 登录存在账号 API返回值
// alibaba.cgame.mp.mpproject.loginexistaccount
//
// 发送消息给游戏
type AlibabaCgameMpMpprojectLoginexistaccountAPIResponse struct {
	model.CommonResponse
	AlibabaCgameMpMpprojectLoginexistaccountAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCgameMpMpprojectLoginexistaccountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCgameMpMpprojectLoginexistaccountAPIResponseModel).Reset()
}

// AlibabaCgameMpMpprojectLoginexistaccountAPIResponseModel is 登录存在账号 成功返回结果
type AlibabaCgameMpMpprojectLoginexistaccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cgame_mp_mpproject_loginexistaccount_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaCgameMpMpprojectLoginexistaccountResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCgameMpMpprojectLoginexistaccountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCgameMpMpprojectLoginexistaccountAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCgameMpMpprojectLoginexistaccountAPIResponse)
	},
}

// GetAlibabaCgameMpMpprojectLoginexistaccountAPIResponse 从 sync.Pool 获取 AlibabaCgameMpMpprojectLoginexistaccountAPIResponse
func GetAlibabaCgameMpMpprojectLoginexistaccountAPIResponse() *AlibabaCgameMpMpprojectLoginexistaccountAPIResponse {
	return poolAlibabaCgameMpMpprojectLoginexistaccountAPIResponse.Get().(*AlibabaCgameMpMpprojectLoginexistaccountAPIResponse)
}

// ReleaseAlibabaCgameMpMpprojectLoginexistaccountAPIResponse 将 AlibabaCgameMpMpprojectLoginexistaccountAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCgameMpMpprojectLoginexistaccountAPIResponse(v *AlibabaCgameMpMpprojectLoginexistaccountAPIResponse) {
	v.Reset()
	poolAlibabaCgameMpMpprojectLoginexistaccountAPIResponse.Put(v)
}
