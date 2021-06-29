package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户有权限的菜单树 API返回值 
alibaba.campus.acl.new.listusermenu

查询用户有权限的菜单树
*/
type AlibabaCampusAclNewListusermenuAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclNewListusermenuResponse
}

// 查询用户有权限的菜单树 成功返回结果
type AlibabaCampusAclNewListusermenuResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_new_listusermenu_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}
