package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewCheckuserpermissionAPIResponse
校验用户是否有权限 API返回值
alibaba.campus.acl.new.checkuserpermission

校验用户是否有权限 */
type AlibabaCampusAclNewCheckuserpermissionAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewCheckuserpermissionAPIResponseModel
}

// AlibabaCampusAclNewCheckuserpermissionAPIResponseModel is 校验用户是否有权限 成功返回结果
type AlibabaCampusAclNewCheckuserpermissionAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_checkuserpermission_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
