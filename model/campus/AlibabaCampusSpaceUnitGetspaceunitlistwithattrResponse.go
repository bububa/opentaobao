package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
空间单元列表带业务属性实例 APIResponse
alibaba.campus.space.unit.getspaceunitlistwithattr

空间单元列表带业务属性实例
*/
type AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceUnitGetspaceunitlistwithattrResponse
}

type AlibabaCampusSpaceUnitGetspaceunitlistwithattrResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_unit_getspaceunitlistwithattr_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
