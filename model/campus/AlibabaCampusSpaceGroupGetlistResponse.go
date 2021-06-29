package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
多条件查询空间分组信息 APIResponse
alibaba.campus.space.group.getlist

多条件查询空间分组信息
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
HSF方法名称：getList
*/
type AlibabaCampusSpaceGroupGetlistAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceGroupGetlistResponse
}

type AlibabaCampusSpaceGroupGetlistResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_group_getlist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
