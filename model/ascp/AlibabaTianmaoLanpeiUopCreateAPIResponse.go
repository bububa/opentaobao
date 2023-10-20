package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoLanpeiUopCreateAPIResponse 阿里巴巴.天猫家装.揽配.履约订单.创建 API返回值
// alibaba.tianmao.lanpei.uop.create
//
// 阿里巴巴.天猫家装.揽配.履约订单.创建
type AlibabaTianmaoLanpeiUopCreateAPIResponse struct {
	model.CommonResponse
	AlibabaTianmaoLanpeiUopCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTianmaoLanpeiUopCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTianmaoLanpeiUopCreateAPIResponseModel).Reset()
}

// AlibabaTianmaoLanpeiUopCreateAPIResponseModel is 阿里巴巴.天猫家装.揽配.履约订单.创建 成功返回结果
type AlibabaTianmaoLanpeiUopCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tianmao_lanpei_uop_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTianmaoLanpeiUopCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTianmaoLanpeiUopCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTianmaoLanpeiUopCreateAPIResponse)
	},
}

// GetAlibabaTianmaoLanpeiUopCreateAPIResponse 从 sync.Pool 获取 AlibabaTianmaoLanpeiUopCreateAPIResponse
func GetAlibabaTianmaoLanpeiUopCreateAPIResponse() *AlibabaTianmaoLanpeiUopCreateAPIResponse {
	return poolAlibabaTianmaoLanpeiUopCreateAPIResponse.Get().(*AlibabaTianmaoLanpeiUopCreateAPIResponse)
}

// ReleaseAlibabaTianmaoLanpeiUopCreateAPIResponse 将 AlibabaTianmaoLanpeiUopCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTianmaoLanpeiUopCreateAPIResponse(v *AlibabaTianmaoLanpeiUopCreateAPIResponse) {
	v.Reset()
	poolAlibabaTianmaoLanpeiUopCreateAPIResponse.Put(v)
}
