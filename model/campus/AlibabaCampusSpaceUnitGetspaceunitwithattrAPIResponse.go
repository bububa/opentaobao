package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse 空间单元id查业务属性实例 API返回值
// alibaba.campus.space.unit.getspaceunitwithattr
//
// 空间单元id查业务属性实例
type AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponseModel).Reset()
}

// AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponseModel is 空间单元id查业务属性实例 成功返回结果
type AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_unit_getspaceunitwithattr_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse)
	},
}

// GetAlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse
func GetAlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse() *AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse {
	return poolAlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse.Get().(*AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse)
}

// ReleaseAlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse 将 AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse(v *AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse.Put(v)
}
