package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcWritesaleslipAPIResponse 开票占库 API返回值
// alibaba.mj.oc.writesaleslip
//
// 开票占库
type AlibabaMjOcWritesaleslipAPIResponse struct {
	model.CommonResponse
	AlibabaMjOcWritesaleslipAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMjOcWritesaleslipAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMjOcWritesaleslipAPIResponseModel).Reset()
}

// AlibabaMjOcWritesaleslipAPIResponseModel is 开票占库 成功返回结果
type AlibabaMjOcWritesaleslipAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_oc_writesaleslip_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMjOcWritesaleslipAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaMjOcWritesaleslipAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMjOcWritesaleslipAPIResponse)
	},
}

// GetAlibabaMjOcWritesaleslipAPIResponse 从 sync.Pool 获取 AlibabaMjOcWritesaleslipAPIResponse
func GetAlibabaMjOcWritesaleslipAPIResponse() *AlibabaMjOcWritesaleslipAPIResponse {
	return poolAlibabaMjOcWritesaleslipAPIResponse.Get().(*AlibabaMjOcWritesaleslipAPIResponse)
}

// ReleaseAlibabaMjOcWritesaleslipAPIResponse 将 AlibabaMjOcWritesaleslipAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMjOcWritesaleslipAPIResponse(v *AlibabaMjOcWritesaleslipAPIResponse) {
	v.Reset()
	poolAlibabaMjOcWritesaleslipAPIResponse.Put(v)
}
