package cloudgame

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCgameMpMpprojectInitnewprojectAPIResponse 创建新的mpproject API返回值
// alibaba.cgame.mp.mpproject.initnewproject
//
// 发送消息给游戏
type AlibabaCgameMpMpprojectInitnewprojectAPIResponse struct {
	model.CommonResponse
	AlibabaCgameMpMpprojectInitnewprojectAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCgameMpMpprojectInitnewprojectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCgameMpMpprojectInitnewprojectAPIResponseModel).Reset()
}

// AlibabaCgameMpMpprojectInitnewprojectAPIResponseModel is 创建新的mpproject 成功返回结果
type AlibabaCgameMpMpprojectInitnewprojectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cgame_mp_mpproject_initnewproject_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaCgameMpMpprojectInitnewprojectResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCgameMpMpprojectInitnewprojectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCgameMpMpprojectInitnewprojectAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCgameMpMpprojectInitnewprojectAPIResponse)
	},
}

// GetAlibabaCgameMpMpprojectInitnewprojectAPIResponse 从 sync.Pool 获取 AlibabaCgameMpMpprojectInitnewprojectAPIResponse
func GetAlibabaCgameMpMpprojectInitnewprojectAPIResponse() *AlibabaCgameMpMpprojectInitnewprojectAPIResponse {
	return poolAlibabaCgameMpMpprojectInitnewprojectAPIResponse.Get().(*AlibabaCgameMpMpprojectInitnewprojectAPIResponse)
}

// ReleaseAlibabaCgameMpMpprojectInitnewprojectAPIResponse 将 AlibabaCgameMpMpprojectInitnewprojectAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCgameMpMpprojectInitnewprojectAPIResponse(v *AlibabaCgameMpMpprojectInitnewprojectAPIResponse) {
	v.Reset()
	poolAlibabaCgameMpMpprojectInitnewprojectAPIResponse.Put(v)
}
