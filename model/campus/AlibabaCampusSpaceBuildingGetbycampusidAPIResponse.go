package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspacebuildinggetbycampusidAPIResponse 根据园区ID获取楼宇 API返回值
// alibaba.campus.space.building.getbycampusid
//
// 根据园区ID获取楼宇
// HSF接口名称：com.alibaba.campus.api.space.service.top.BuildingApiTopService
// HSF方法名称：getBuildingList
type AlibabacampusspacebuildinggetbycampusidAPIResponse struct {
	model.CommonResponse
	AlibabacampusspacebuildinggetbycampusidAPIResponseModel
}

// AlibabacampusspacebuildinggetbycampusidAPIResponseModel is 根据园区ID获取楼宇 成功返回结果
type AlibabacampusspacebuildinggetbycampusidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_building_getbycampusid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}
