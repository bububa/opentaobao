package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenDeletefloorAPIResponse 大麦换验平台-第三方对外开放-楼层接口deleteFloor API返回值
// alibaba.damai.mev.open.deletefloor
//
// deleteFloor
type AlibabaDamaiMevOpenDeletefloorAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenDeletefloorAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenDeletefloorAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenDeletefloorAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenDeletefloorAPIResponseModel is 大麦换验平台-第三方对外开放-楼层接口deleteFloor 成功返回结果
type AlibabaDamaiMevOpenDeletefloorAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_deletefloor_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenDeletefloorResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenDeletefloorAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenDeletefloorAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenDeletefloorAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenDeletefloorAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenDeletefloorAPIResponse
func GetAlibabaDamaiMevOpenDeletefloorAPIResponse() *AlibabaDamaiMevOpenDeletefloorAPIResponse {
	return poolAlibabaDamaiMevOpenDeletefloorAPIResponse.Get().(*AlibabaDamaiMevOpenDeletefloorAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenDeletefloorAPIResponse 将 AlibabaDamaiMevOpenDeletefloorAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenDeletefloorAPIResponse(v *AlibabaDamaiMevOpenDeletefloorAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenDeletefloorAPIResponse.Put(v)
}
