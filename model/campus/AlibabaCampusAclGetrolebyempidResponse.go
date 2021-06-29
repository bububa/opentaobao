package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据用户查询角色 APIResponse
alibaba.campus.acl.getrolebyempid

根据用户查询角色
*/
type AlibabaCampusAclGetrolebyempidAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclGetrolebyempidResponse
}

type AlibabaCampusAclGetrolebyempidResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_getrolebyempid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *CollectionResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
