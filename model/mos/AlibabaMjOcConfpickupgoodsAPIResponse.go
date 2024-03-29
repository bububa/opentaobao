package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcConfpickupgoodsAPIResponse 提货核销 API返回值
// alibaba.mj.oc.confpickupgoods
//
// 此API用于在银泰商场中，消费者在提货中心提货时， 商户后台调用此接口进行提货核销操作
type AlibabaMjOcConfpickupgoodsAPIResponse struct {
	model.CommonResponse
	AlibabaMjOcConfpickupgoodsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMjOcConfpickupgoodsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMjOcConfpickupgoodsAPIResponseModel).Reset()
}

// AlibabaMjOcConfpickupgoodsAPIResponseModel is 提货核销 成功返回结果
type AlibabaMjOcConfpickupgoodsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_oc_confpickupgoods_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMjOcConfpickupgoodsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaMjOcConfpickupgoodsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMjOcConfpickupgoodsAPIResponse)
	},
}

// GetAlibabaMjOcConfpickupgoodsAPIResponse 从 sync.Pool 获取 AlibabaMjOcConfpickupgoodsAPIResponse
func GetAlibabaMjOcConfpickupgoodsAPIResponse() *AlibabaMjOcConfpickupgoodsAPIResponse {
	return poolAlibabaMjOcConfpickupgoodsAPIResponse.Get().(*AlibabaMjOcConfpickupgoodsAPIResponse)
}

// ReleaseAlibabaMjOcConfpickupgoodsAPIResponse 将 AlibabaMjOcConfpickupgoodsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMjOcConfpickupgoodsAPIResponse(v *AlibabaMjOcConfpickupgoodsAPIResponse) {
	v.Reset()
	poolAlibabaMjOcConfpickupgoodsAPIResponse.Put(v)
}
