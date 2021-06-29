package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据id获取楼层 API返回值 
alibaba.campus.space.floor.getbyid

根据id获取楼层
*/
type AlibabaCampusSpaceFloorGetbyidAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceFloorGetbyidResponse
}

// 根据id获取楼层 成功返回结果
type AlibabaCampusSpaceFloorGetbyidResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_floor_getbyid_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
