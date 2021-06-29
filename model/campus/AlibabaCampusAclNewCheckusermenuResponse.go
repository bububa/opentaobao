package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否有菜单权限 APIResponse
alibaba.campus.acl.new.checkusermenu

校验用户是否有菜单权限
*/
type AlibabaCampusAclNewCheckusermenuAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclNewCheckusermenuResponse
}

type AlibabaCampusAclNewCheckusermenuResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_new_checkusermenu_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
