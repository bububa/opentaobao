package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyBillDetailQueryAPIResponse 账单明细接口 API返回值
// alibaba.tcls.aelophy.bill.detail.query
//
// 账单明细接口
type AlibabaTclsAelophyBillDetailQueryAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyBillDetailQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyBillDetailQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyBillDetailQueryAPIResponseModel).Reset()
}

// AlibabaTclsAelophyBillDetailQueryAPIResponseModel is 账单明细接口 成功返回结果
type AlibabaTclsAelophyBillDetailQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_bill_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	ApiResult *ApiPageResults `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyBillDetailQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaTclsAelophyBillDetailQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyBillDetailQueryAPIResponse)
	},
}

// GetAlibabaTclsAelophyBillDetailQueryAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyBillDetailQueryAPIResponse
func GetAlibabaTclsAelophyBillDetailQueryAPIResponse() *AlibabaTclsAelophyBillDetailQueryAPIResponse {
	return poolAlibabaTclsAelophyBillDetailQueryAPIResponse.Get().(*AlibabaTclsAelophyBillDetailQueryAPIResponse)
}

// ReleaseAlibabaTclsAelophyBillDetailQueryAPIResponse 将 AlibabaTclsAelophyBillDetailQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyBillDetailQueryAPIResponse(v *AlibabaTclsAelophyBillDetailQueryAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyBillDetailQueryAPIResponse.Put(v)
}
