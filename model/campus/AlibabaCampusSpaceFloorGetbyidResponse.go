package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据id获取楼层 APIResponse
alibaba.campus.space.floor.getbyid

根据id获取楼层
*/
type AlibabaCampusSpaceFloorGetbyidAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceFloorGetbyidResponse
}

type AlibabaCampusSpaceFloorGetbyidResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_floor_getbyid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
