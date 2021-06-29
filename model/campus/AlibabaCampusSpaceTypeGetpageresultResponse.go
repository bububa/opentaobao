package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询空间类别接口 APIResponse
alibaba.campus.space.type.getpageresult

分页查询空间类别接口
HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
HSF方法名称：getPageResult
*/
type AlibabaCampusSpaceTypeGetpageresultAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceTypeGetpageresultResponse
}

type AlibabaCampusSpaceTypeGetpageresultResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_type_getpageresult_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
