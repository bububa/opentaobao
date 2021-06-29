package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】 APIResponse
alibaba.campus.space.unit.getlistwithattrbygroupid

根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】
*/
type AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceUnitGetlistwithattrbygroupidResponse
}

type AlibabaCampusSpaceUnitGetlistwithattrbygroupidResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_unit_getlistwithattrbygroupid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
