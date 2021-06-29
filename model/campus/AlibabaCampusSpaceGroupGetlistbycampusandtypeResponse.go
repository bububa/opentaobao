package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据园区id及TypeId获取空间分组 APIResponse
alibaba.campus.space.group.getlistbycampusandtype

根据园区id及TypeId获取空间分组
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
HSF方法名称：getListByCampusAndType
*/
type AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceGroupGetlistbycampusandtypeResponse
}

type AlibabaCampusSpaceGroupGetlistbycampusandtypeResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_group_getlistbycampusandtype_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
