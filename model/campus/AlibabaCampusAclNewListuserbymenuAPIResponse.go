package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewListuserbymenuAPIResponse 查询菜单下的人员 API返回值
// alibaba.campus.acl.new.listuserbymenu
//
// 查询拥有菜单权限的用户
type AlibabaCampusAclNewListuserbymenuAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewListuserbymenuAPIResponseModel
}

// AlibabaCampusAclNewListuserbymenuAPIResponseModel is 查询菜单下的人员 成功返回结果
type AlibabaCampusAclNewListuserbymenuAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_listuserbymenu_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}
