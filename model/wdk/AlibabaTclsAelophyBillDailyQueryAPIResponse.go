package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyBillDailyQueryAPIResponse 账单日汇总接口 API返回值
// alibaba.tcls.aelophy.bill.daily.query
//
// 账单日汇总接口
type AlibabaTclsAelophyBillDailyQueryAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyBillDailyQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyBillDailyQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyBillDailyQueryAPIResponseModel).Reset()
}

// AlibabaTclsAelophyBillDailyQueryAPIResponseModel is 账单日汇总接口 成功返回结果
type AlibabaTclsAelophyBillDailyQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_bill_daily_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	ApiResult *ApiPageResults `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyBillDailyQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaTclsAelophyBillDailyQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyBillDailyQueryAPIResponse)
	},
}

// GetAlibabaTclsAelophyBillDailyQueryAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyBillDailyQueryAPIResponse
func GetAlibabaTclsAelophyBillDailyQueryAPIResponse() *AlibabaTclsAelophyBillDailyQueryAPIResponse {
	return poolAlibabaTclsAelophyBillDailyQueryAPIResponse.Get().(*AlibabaTclsAelophyBillDailyQueryAPIResponse)
}

// ReleaseAlibabaTclsAelophyBillDailyQueryAPIResponse 将 AlibabaTclsAelophyBillDailyQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyBillDailyQueryAPIResponse(v *AlibabaTclsAelophyBillDailyQueryAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyBillDailyQueryAPIResponse.Put(v)
}
