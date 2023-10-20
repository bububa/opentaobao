package mydata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMydataSelfProductGetAPIResponse 获取客户产品相关表现数据 API返回值
// alibaba.mydata.self.product.get
//
// 获取客户产品相关表现数据
type AlibabaMydataSelfProductGetAPIResponse struct {
	model.CommonResponse
	AlibabaMydataSelfProductGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMydataSelfProductGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMydataSelfProductGetAPIResponseModel).Reset()
}

// AlibabaMydataSelfProductGetAPIResponseModel is 获取客户产品相关表现数据 成功返回结果
type AlibabaMydataSelfProductGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mydata_self_product_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品效果查询结果
	Result *ProductEffect `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMydataSelfProductGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMydataSelfProductGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMydataSelfProductGetAPIResponse)
	},
}

// GetAlibabaMydataSelfProductGetAPIResponse 从 sync.Pool 获取 AlibabaMydataSelfProductGetAPIResponse
func GetAlibabaMydataSelfProductGetAPIResponse() *AlibabaMydataSelfProductGetAPIResponse {
	return poolAlibabaMydataSelfProductGetAPIResponse.Get().(*AlibabaMydataSelfProductGetAPIResponse)
}

// ReleaseAlibabaMydataSelfProductGetAPIResponse 将 AlibabaMydataSelfProductGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMydataSelfProductGetAPIResponse(v *AlibabaMydataSelfProductGetAPIResponse) {
	v.Reset()
	poolAlibabaMydataSelfProductGetAPIResponse.Put(v)
}
