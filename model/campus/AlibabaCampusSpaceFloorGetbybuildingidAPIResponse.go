package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceFloorGetbybuildingidAPIResponse 根据楼宇ID获取楼层 API返回值
// alibaba.campus.space.floor.getbybuildingid
//
// 根据楼宇ID获取楼层
// HSF接口名称：com.alibaba.campus.api.space.service.top.FloorApiTopService
// HSF方法名称：getFloorList
type AlibabaCampusSpaceFloorGetbybuildingidAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceFloorGetbybuildingidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceFloorGetbybuildingidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceFloorGetbybuildingidAPIResponseModel).Reset()
}

// AlibabaCampusSpaceFloorGetbybuildingidAPIResponseModel is 根据楼宇ID获取楼层 成功返回结果
type AlibabaCampusSpaceFloorGetbybuildingidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_floor_getbybuildingid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceFloorGetbybuildingidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceFloorGetbybuildingidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceFloorGetbybuildingidAPIResponse)
	},
}

// GetAlibabaCampusSpaceFloorGetbybuildingidAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceFloorGetbybuildingidAPIResponse
func GetAlibabaCampusSpaceFloorGetbybuildingidAPIResponse() *AlibabaCampusSpaceFloorGetbybuildingidAPIResponse {
	return poolAlibabaCampusSpaceFloorGetbybuildingidAPIResponse.Get().(*AlibabaCampusSpaceFloorGetbybuildingidAPIResponse)
}

// ReleaseAlibabaCampusSpaceFloorGetbybuildingidAPIResponse 将 AlibabaCampusSpaceFloorGetbybuildingidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceFloorGetbybuildingidAPIResponse(v *AlibabaCampusSpaceFloorGetbybuildingidAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceFloorGetbybuildingidAPIResponse.Put(v)
}
