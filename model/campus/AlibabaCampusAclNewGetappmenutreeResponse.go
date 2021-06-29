package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询应用下的菜单树 APIResponse
alibaba.campus.acl.new.getappmenutree

查询应用下的菜单树
*/
type AlibabaCampusAclNewGetappmenutreeAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclNewGetappmenutreeResponse
}

type AlibabaCampusAclNewGetappmenutreeResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_new_getappmenutree_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回结果
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
