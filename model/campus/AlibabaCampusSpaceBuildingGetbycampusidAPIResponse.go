package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceBuildingGetbycampusidAPIResponse 根据园区ID获取楼宇 API返回值
// alibaba.campus.space.building.getbycampusid
//
// 根据园区ID获取楼宇
// HSF接口名称：com.alibaba.campus.api.space.service.top.BuildingApiTopService
// HSF方法名称：getBuildingList
type AlibabaCampusSpaceBuildingGetbycampusidAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceBuildingGetbycampusidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceBuildingGetbycampusidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceBuildingGetbycampusidAPIResponseModel).Reset()
}

// AlibabaCampusSpaceBuildingGetbycampusidAPIResponseModel is 根据园区ID获取楼宇 成功返回结果
type AlibabaCampusSpaceBuildingGetbycampusidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_building_getbycampusid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceBuildingGetbycampusidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceBuildingGetbycampusidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceBuildingGetbycampusidAPIResponse)
	},
}

// GetAlibabaCampusSpaceBuildingGetbycampusidAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceBuildingGetbycampusidAPIResponse
func GetAlibabaCampusSpaceBuildingGetbycampusidAPIResponse() *AlibabaCampusSpaceBuildingGetbycampusidAPIResponse {
	return poolAlibabaCampusSpaceBuildingGetbycampusidAPIResponse.Get().(*AlibabaCampusSpaceBuildingGetbycampusidAPIResponse)
}

// ReleaseAlibabaCampusSpaceBuildingGetbycampusidAPIResponse 将 AlibabaCampusSpaceBuildingGetbycampusidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceBuildingGetbycampusidAPIResponse(v *AlibabaCampusSpaceBuildingGetbycampusidAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceBuildingGetbycampusidAPIResponse.Put(v)
}
