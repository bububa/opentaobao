package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据分组ID查询相应的空间单元 APIResponse
alibaba.campus.space.unit.getlistbygroupid

根据分组ID查询相应的空间单元
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：getListByGroupId
*/
type AlibabaCampusSpaceUnitGetlistbygroupidAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceUnitGetlistbygroupidResponse
}

type AlibabaCampusSpaceUnitGetlistbygroupidResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_unit_getlistbygroupid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
