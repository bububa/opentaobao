package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据类别编码查询类别 APIResponse
alibaba.campus.space.type.getbycode

根据类别编码查询类别
HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
HSF方法名称：getByCode
*/
type AlibabaCampusSpaceTypeGetbycodeAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceTypeGetbycodeResponse
}

type AlibabaCampusSpaceTypeGetbycodeResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_type_getbycode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
