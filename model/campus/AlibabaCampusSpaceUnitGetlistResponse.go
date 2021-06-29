package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
多条件查询空间单元信息 APIResponse
alibaba.campus.space.unit.getlist

多条件查询空间单元信息
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：getList
*/
type AlibabaCampusSpaceUnitGetlistAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceUnitGetlistResponse
}

type AlibabaCampusSpaceUnitGetlistResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_unit_getlist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // results
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
