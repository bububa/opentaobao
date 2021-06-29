package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询员工全部权限(包括角色下面的权限) APIResponse
alibaba.campus.acl.queryallemppermiitem

查询员工全部权限(包括角色下面的权限)
*/
type AlibabaCampusAclQueryallemppermiitemAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclQueryallemppermiitemResponse
}

type AlibabaCampusAclQueryallemppermiitemResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_queryallemppermiitem_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
