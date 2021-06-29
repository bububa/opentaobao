package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据园区id获取园区信息 APIResponse
alibaba.campus.space.campus.getbyid

根据园区id获取园区信息
HSF接口名称：com.alibaba.campus.api.space.service.top.CampusApiTopService
HSF方法名称：getCampusById
*/
type AlibabaCampusSpaceCampusGetbyidAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceCampusGetbyidResponse
}

type AlibabaCampusSpaceCampusGetbyidResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_campus_getbyid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
