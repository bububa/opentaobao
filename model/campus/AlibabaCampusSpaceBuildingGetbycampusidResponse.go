package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据园区ID获取楼宇 APIResponse
alibaba.campus.space.building.getbycampusid

根据园区ID获取楼宇
HSF接口名称：com.alibaba.campus.api.space.service.top.BuildingApiTopService
HSF方法名称：getBuildingList
*/
type AlibabaCampusSpaceBuildingGetbycampusidAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceBuildingGetbycampusidResponse
}

type AlibabaCampusSpaceBuildingGetbycampusidResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_building_getbycampusid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
