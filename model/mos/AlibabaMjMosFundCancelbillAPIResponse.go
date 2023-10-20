package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjMosFundCancelbillAPIResponse 取消付款单 API返回值
// alibaba.mj.mos.fund.cancelbill
//
// 取消付款单
type AlibabaMjMosFundCancelbillAPIResponse struct {
	model.CommonResponse
	AlibabaMjMosFundCancelbillAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMjMosFundCancelbillAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMjMosFundCancelbillAPIResponseModel).Reset()
}

// AlibabaMjMosFundCancelbillAPIResponseModel is 取消付款单 成功返回结果
type AlibabaMjMosFundCancelbillAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_mos_fund_cancelbill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMjMosFundCancelbillAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = false
}

var poolAlibabaMjMosFundCancelbillAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMjMosFundCancelbillAPIResponse)
	},
}

// GetAlibabaMjMosFundCancelbillAPIResponse 从 sync.Pool 获取 AlibabaMjMosFundCancelbillAPIResponse
func GetAlibabaMjMosFundCancelbillAPIResponse() *AlibabaMjMosFundCancelbillAPIResponse {
	return poolAlibabaMjMosFundCancelbillAPIResponse.Get().(*AlibabaMjMosFundCancelbillAPIResponse)
}

// ReleaseAlibabaMjMosFundCancelbillAPIResponse 将 AlibabaMjMosFundCancelbillAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMjMosFundCancelbillAPIResponse(v *AlibabaMjMosFundCancelbillAPIResponse) {
	v.Reset()
	poolAlibabaMjMosFundCancelbillAPIResponse.Put(v)
}
