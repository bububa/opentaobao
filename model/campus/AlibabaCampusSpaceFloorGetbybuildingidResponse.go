package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据楼宇ID获取楼层 API返回值 
alibaba.campus.space.floor.getbybuildingid

根据楼宇ID获取楼层
HSF接口名称：com.alibaba.campus.api.space.service.top.FloorApiTopService
HSF方法名称：getFloorList
*/
type AlibabaCampusSpaceFloorGetbybuildingidAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceFloorGetbybuildingidResponse
}

// 根据楼宇ID获取楼层 成功返回结果
type AlibabaCampusSpaceFloorGetbybuildingidResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_floor_getbybuildingid_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}
