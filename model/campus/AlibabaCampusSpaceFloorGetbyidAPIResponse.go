package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceFloorGetbyidAPIResponse 根据id获取楼层 API返回值
// alibaba.campus.space.floor.getbyid
//
// 根据id获取楼层
type AlibabaCampusSpaceFloorGetbyidAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceFloorGetbyidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceFloorGetbyidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceFloorGetbyidAPIResponseModel).Reset()
}

// AlibabaCampusSpaceFloorGetbyidAPIResponseModel is 根据id获取楼层 成功返回结果
type AlibabaCampusSpaceFloorGetbyidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_floor_getbyid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceFloorGetbyidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceFloorGetbyidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceFloorGetbyidAPIResponse)
	},
}

// GetAlibabaCampusSpaceFloorGetbyidAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceFloorGetbyidAPIResponse
func GetAlibabaCampusSpaceFloorGetbyidAPIResponse() *AlibabaCampusSpaceFloorGetbyidAPIResponse {
	return poolAlibabaCampusSpaceFloorGetbyidAPIResponse.Get().(*AlibabaCampusSpaceFloorGetbyidAPIResponse)
}

// ReleaseAlibabaCampusSpaceFloorGetbyidAPIResponse 将 AlibabaCampusSpaceFloorGetbyidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceFloorGetbyidAPIResponse(v *AlibabaCampusSpaceFloorGetbyidAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceFloorGetbyidAPIResponse.Put(v)
}
