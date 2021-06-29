package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否有权限 APIResponse
alibaba.campus.acl.new.checkuserpermission

校验用户是否有权限
*/
type AlibabaCampusAclNewCheckuserpermissionAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclNewCheckuserpermissionResponse
}

type AlibabaCampusAclNewCheckuserpermissionResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_new_checkuserpermission_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
