package mydata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMydataSelfProductDateGetAPIResponse 获取客户产品相关表现数据的可用时间范围 API返回值
// alibaba.mydata.self.product.date.get
//
// 获取客户产品相关表现数据的可用时间范围
type AlibabaMydataSelfProductDateGetAPIResponse struct {
	model.CommonResponse
	AlibabaMydataSelfProductDateGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMydataSelfProductDateGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMydataSelfProductDateGetAPIResponseModel).Reset()
}

// AlibabaMydataSelfProductDateGetAPIResponseModel is 获取客户产品相关表现数据的可用时间范围 成功返回结果
type AlibabaMydataSelfProductDateGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mydata_self_product_date_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// endDate
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// startDate
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMydataSelfProductDateGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EndDate = ""
	m.StartDate = ""
}

var poolAlibabaMydataSelfProductDateGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMydataSelfProductDateGetAPIResponse)
	},
}

// GetAlibabaMydataSelfProductDateGetAPIResponse 从 sync.Pool 获取 AlibabaMydataSelfProductDateGetAPIResponse
func GetAlibabaMydataSelfProductDateGetAPIResponse() *AlibabaMydataSelfProductDateGetAPIResponse {
	return poolAlibabaMydataSelfProductDateGetAPIResponse.Get().(*AlibabaMydataSelfProductDateGetAPIResponse)
}

// ReleaseAlibabaMydataSelfProductDateGetAPIResponse 将 AlibabaMydataSelfProductDateGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMydataSelfProductDateGetAPIResponse(v *AlibabaMydataSelfProductDateGetAPIResponse) {
	v.Reset()
	poolAlibabaMydataSelfProductDateGetAPIResponse.Put(v)
}
