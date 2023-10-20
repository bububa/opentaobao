package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemmappingUnbundleAPIResponse 商货关联解绑 API返回值
// alibaba.dchain.aoxiang.itemmapping.unbundle
//
// 商货关联解绑
type AlibabaDchainAoxiangItemmappingUnbundleAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangItemmappingUnbundleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangItemmappingUnbundleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangItemmappingUnbundleAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangItemmappingUnbundleAPIResponseModel is 商货关联解绑 成功返回结果
type AlibabaDchainAoxiangItemmappingUnbundleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_itemmapping_unbundle_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	UnbundleItemmappingResponse *UnbundleItemmappingResponse `json:"unbundle_itemmapping_response,omitempty" xml:"unbundle_itemmapping_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangItemmappingUnbundleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UnbundleItemmappingResponse = nil
}

var poolAlibabaDchainAoxiangItemmappingUnbundleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangItemmappingUnbundleAPIResponse)
	},
}

// GetAlibabaDchainAoxiangItemmappingUnbundleAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangItemmappingUnbundleAPIResponse
func GetAlibabaDchainAoxiangItemmappingUnbundleAPIResponse() *AlibabaDchainAoxiangItemmappingUnbundleAPIResponse {
	return poolAlibabaDchainAoxiangItemmappingUnbundleAPIResponse.Get().(*AlibabaDchainAoxiangItemmappingUnbundleAPIResponse)
}

// ReleaseAlibabaDchainAoxiangItemmappingUnbundleAPIResponse 将 AlibabaDchainAoxiangItemmappingUnbundleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangItemmappingUnbundleAPIResponse(v *AlibabaDchainAoxiangItemmappingUnbundleAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangItemmappingUnbundleAPIResponse.Put(v)
}
