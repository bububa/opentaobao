package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户有权限的菜单树 APIResponse
alibaba.campus.acl.new.listusermenu

查询用户有权限的菜单树
*/
type AlibabaCampusAclNewListusermenuAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclNewListusermenuResponse
}

type AlibabaCampusAclNewListusermenuResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_new_listusermenu_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
