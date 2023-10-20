package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripEmployeeQueryAPIResponse 企业员工查询 API返回值
// alitrip.btrip.employee.query
//
// 企业员工查询
type AlitripBtripEmployeeQueryAPIResponse struct {
	model.CommonResponse
	AlitripBtripEmployeeQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripEmployeeQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripEmployeeQueryAPIResponseModel).Reset()
}

// AlitripBtripEmployeeQueryAPIResponseModel is 企业员工查询 成功返回结果
type AlitripBtripEmployeeQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_employee_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象。
	Result *BtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripEmployeeQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripEmployeeQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripEmployeeQueryAPIResponse)
	},
}

// GetAlitripBtripEmployeeQueryAPIResponse 从 sync.Pool 获取 AlitripBtripEmployeeQueryAPIResponse
func GetAlitripBtripEmployeeQueryAPIResponse() *AlitripBtripEmployeeQueryAPIResponse {
	return poolAlitripBtripEmployeeQueryAPIResponse.Get().(*AlitripBtripEmployeeQueryAPIResponse)
}

// ReleaseAlitripBtripEmployeeQueryAPIResponse 将 AlitripBtripEmployeeQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripEmployeeQueryAPIResponse(v *AlitripBtripEmployeeQueryAPIResponse) {
	v.Reset()
	poolAlitripBtripEmployeeQueryAPIResponse.Put(v)
}
