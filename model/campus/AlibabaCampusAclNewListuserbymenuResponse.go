package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询菜单下的人员 APIResponse
alibaba.campus.acl.new.listuserbymenu

查询拥有菜单权限的用户
*/
type AlibabaCampusAclNewListuserbymenuAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclNewListuserbymenuResponse
}

type AlibabaCampusAclNewListuserbymenuResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_new_listuserbymenu_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
