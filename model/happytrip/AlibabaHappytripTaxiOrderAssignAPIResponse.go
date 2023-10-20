package happytrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiOrderAssignAPIResponse 订单指派 API返回值
// alibaba.happytrip.taxi.order.assign
//
// 通知供应商订单指派成功
type AlibabaHappytripTaxiOrderAssignAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiOrderAssignAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiOrderAssignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHappytripTaxiOrderAssignAPIResponseModel).Reset()
}

// AlibabaHappytripTaxiOrderAssignAPIResponseModel is 订单指派 成功返回结果
type AlibabaHappytripTaxiOrderAssignAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_order_assign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误描述
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 错误代码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiOrderAssignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errmsg = ""
	m.Errno = 0
}

var poolAlibabaHappytripTaxiOrderAssignAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripTaxiOrderAssignAPIResponse)
	},
}

// GetAlibabaHappytripTaxiOrderAssignAPIResponse 从 sync.Pool 获取 AlibabaHappytripTaxiOrderAssignAPIResponse
func GetAlibabaHappytripTaxiOrderAssignAPIResponse() *AlibabaHappytripTaxiOrderAssignAPIResponse {
	return poolAlibabaHappytripTaxiOrderAssignAPIResponse.Get().(*AlibabaHappytripTaxiOrderAssignAPIResponse)
}

// ReleaseAlibabaHappytripTaxiOrderAssignAPIResponse 将 AlibabaHappytripTaxiOrderAssignAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHappytripTaxiOrderAssignAPIResponse(v *AlibabaHappytripTaxiOrderAssignAPIResponse) {
	v.Reset()
	poolAlibabaHappytripTaxiOrderAssignAPIResponse.Put(v)
}
