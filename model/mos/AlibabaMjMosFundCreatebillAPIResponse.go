package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjMosFundCreatebillAPIResponse 创建一个付款单 API返回值
// alibaba.mj.mos.fund.createbill
//
// 创建一个付款单
type AlibabaMjMosFundCreatebillAPIResponse struct {
	model.CommonResponse
	AlibabaMjMosFundCreatebillAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMjMosFundCreatebillAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMjMosFundCreatebillAPIResponseModel).Reset()
}

// AlibabaMjMosFundCreatebillAPIResponseModel is 创建一个付款单 成功返回结果
type AlibabaMjMosFundCreatebillAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_mos_fund_createbill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMjMosFundCreatebillAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
}

var poolAlibabaMjMosFundCreatebillAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMjMosFundCreatebillAPIResponse)
	},
}

// GetAlibabaMjMosFundCreatebillAPIResponse 从 sync.Pool 获取 AlibabaMjMosFundCreatebillAPIResponse
func GetAlibabaMjMosFundCreatebillAPIResponse() *AlibabaMjMosFundCreatebillAPIResponse {
	return poolAlibabaMjMosFundCreatebillAPIResponse.Get().(*AlibabaMjMosFundCreatebillAPIResponse)
}

// ReleaseAlibabaMjMosFundCreatebillAPIResponse 将 AlibabaMjMosFundCreatebillAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMjMosFundCreatebillAPIResponse(v *AlibabaMjMosFundCreatebillAPIResponse) {
	v.Reset()
	poolAlibabaMjMosFundCreatebillAPIResponse.Put(v)
}
