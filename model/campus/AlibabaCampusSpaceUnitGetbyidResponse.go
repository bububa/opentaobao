package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据ID查询指定空间单元信息 APIResponse
alibaba.campus.space.unit.getbyid

根据ID查询指定空间单元信息
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：getById
*/
type AlibabaCampusSpaceUnitGetbyidAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceUnitGetbyidResponse
}

type AlibabaCampusSpaceUnitGetbyidResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_unit_getbyid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
