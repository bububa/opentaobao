package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询空间分组业务属性 APIResponse
alibaba.campus.space.group.getspacegrouplistwithattr

分页查询空间分组业务属性
*/
type AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceGroupGetspacegrouplistwithattrResponse
}

type AlibabaCampusSpaceGroupGetspacegrouplistwithattrResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_group_getspacegrouplistwithattr_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
