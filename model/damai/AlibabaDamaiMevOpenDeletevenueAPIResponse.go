package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenDeletevenueAPIResponse 大麦换验平台-第三方对外开放-场馆接口deleteVenue API返回值
// alibaba.damai.mev.open.deletevenue
//
// 开放接口，删除场馆
type AlibabaDamaiMevOpenDeletevenueAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenDeletevenueAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenDeletevenueAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenDeletevenueAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenDeletevenueAPIResponseModel is 大麦换验平台-第三方对外开放-场馆接口deleteVenue 成功返回结果
type AlibabaDamaiMevOpenDeletevenueAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_deletevenue_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenDeletevenueResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenDeletevenueAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenDeletevenueAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenDeletevenueAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenDeletevenueAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenDeletevenueAPIResponse
func GetAlibabaDamaiMevOpenDeletevenueAPIResponse() *AlibabaDamaiMevOpenDeletevenueAPIResponse {
	return poolAlibabaDamaiMevOpenDeletevenueAPIResponse.Get().(*AlibabaDamaiMevOpenDeletevenueAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenDeletevenueAPIResponse 将 AlibabaDamaiMevOpenDeletevenueAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenDeletevenueAPIResponse(v *AlibabaDamaiMevOpenDeletevenueAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenDeletevenueAPIResponse.Put(v)
}
