package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTcCompassWarehousenetworkQueryAPIResponse 按仓维度来查询鸟潮网络 API返回值
// alibaba.tc.compass.warehousenetwork.query
//
// 按仓维度来查询鸟潮网络
type AlibabaTcCompassWarehousenetworkQueryAPIResponse struct {
	model.CommonResponse
	AlibabaTcCompassWarehousenetworkQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTcCompassWarehousenetworkQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTcCompassWarehousenetworkQueryAPIResponseModel).Reset()
}

// AlibabaTcCompassWarehousenetworkQueryAPIResponseModel is 按仓维度来查询鸟潮网络 成功返回结果
type AlibabaTcCompassWarehousenetworkQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tc_compass_warehousenetwork_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTcCompassWarehousenetworkQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTcCompassWarehousenetworkQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTcCompassWarehousenetworkQueryAPIResponse)
	},
}

// GetAlibabaTcCompassWarehousenetworkQueryAPIResponse 从 sync.Pool 获取 AlibabaTcCompassWarehousenetworkQueryAPIResponse
func GetAlibabaTcCompassWarehousenetworkQueryAPIResponse() *AlibabaTcCompassWarehousenetworkQueryAPIResponse {
	return poolAlibabaTcCompassWarehousenetworkQueryAPIResponse.Get().(*AlibabaTcCompassWarehousenetworkQueryAPIResponse)
}

// ReleaseAlibabaTcCompassWarehousenetworkQueryAPIResponse 将 AlibabaTcCompassWarehousenetworkQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTcCompassWarehousenetworkQueryAPIResponse(v *AlibabaTcCompassWarehousenetworkQueryAPIResponse) {
	v.Reset()
	poolAlibabaTcCompassWarehousenetworkQueryAPIResponse.Put(v)
}
