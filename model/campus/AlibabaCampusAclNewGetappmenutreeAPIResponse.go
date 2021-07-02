package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewGetappmenutreeAPIResponse 查询应用下的菜单树 API返回值
// alibaba.campus.acl.new.getappmenutree
//
// 查询应用下的菜单树
type AlibabaCampusAclNewGetappmenutreeAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewGetappmenutreeAPIResponseModel
}

// AlibabaCampusAclNewGetappmenutreeAPIResponseModel is 查询应用下的菜单树 成功返回结果
type AlibabaCampusAclNewGetappmenutreeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_getappmenutree_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}
