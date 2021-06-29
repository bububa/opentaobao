package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增查询多个分组ID各自相关的空间单元信息 APIResponse
alibaba.campus.space.unit.getlistmapbygroupid

新增查询多个分组ID各自相关的空间单元信息
HSF接口名称：	com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：	getListMapByGroupIds
*/
type AlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceUnitGetlistmapbygroupidResponse
}

type AlibabaCampusSpaceUnitGetlistmapbygroupidResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_unit_getlistmapbygroupid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaCampusSpaceUnitGetlistmapbygroupidMapResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
