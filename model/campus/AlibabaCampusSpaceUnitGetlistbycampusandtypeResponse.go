package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据园区id及TypeId获取空间单元 APIResponse
alibaba.campus.space.unit.getlistbycampusandtype

根据园区id及TypeId获取空间单元
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：getListByCampusAndType
*/
type AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceUnitGetlistbycampusandtypeResponse
}

type AlibabaCampusSpaceUnitGetlistbycampusandtypeResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_unit_getlistbycampusandtype_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
