package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse 空间单元列表带业务属性实例 API返回值
// alibaba.campus.space.unit.getspaceunitlistwithattr
//
// 空间单元列表带业务属性实例
type AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponseModel).Reset()
}

// AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponseModel is 空间单元列表带业务属性实例 成功返回结果
type AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_unit_getspaceunitlistwithattr_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse)
	},
}

// GetAlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse
func GetAlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse() *AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse {
	return poolAlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse.Get().(*AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse)
}

// ReleaseAlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse 将 AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse(v *AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse.Put(v)
}
