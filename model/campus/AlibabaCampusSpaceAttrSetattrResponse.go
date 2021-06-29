package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增业务属性实例接口 APIResponse
alibaba.campus.space.attr.setattr

新增业务属性实例接口
*/
type AlibabaCampusSpaceAttrSetattrAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceAttrSetattrResponse
}

type AlibabaCampusSpaceAttrSetattrResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_attr_setattr_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
