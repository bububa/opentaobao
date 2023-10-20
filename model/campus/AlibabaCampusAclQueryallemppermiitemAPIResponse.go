package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclqueryallemppermiitemAPIResponse 查询员工全部权限(包括角色下面的权限) API返回值
// alibaba.campus.acl.queryallemppermiitem
//
// 查询员工全部权限(包括角色下面的权限)
type AlibabacampusaclqueryallemppermiitemAPIResponse struct {
	model.CommonResponse
	AlibabacampusaclqueryallemppermiitemAPIResponseModel
}

// AlibabacampusaclqueryallemppermiitemAPIResponseModel is 查询员工全部权限(包括角色下面的权限) 成功返回结果
type AlibabacampusaclqueryallemppermiitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_queryallemppermiitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}
