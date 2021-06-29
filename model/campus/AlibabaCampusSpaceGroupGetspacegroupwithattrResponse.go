package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
空间分组id查业务属性实例 APIResponse
alibaba.campus.space.group.getspacegroupwithattr

空间分组id查业务属性实例
*/
type AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceGroupGetspacegroupwithattrResponse
}

type AlibabaCampusSpaceGroupGetspacegroupwithattrResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_group_getspacegroupwithattr_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
