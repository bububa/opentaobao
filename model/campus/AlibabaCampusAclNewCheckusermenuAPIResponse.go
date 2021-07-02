package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewCheckusermenuAPIResponse 校验用户是否有菜单权限 API返回值
// alibaba.campus.acl.new.checkusermenu
//
// 校验用户是否有菜单权限
type AlibabaCampusAclNewCheckusermenuAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewCheckusermenuAPIResponseModel
}

// AlibabaCampusAclNewCheckusermenuAPIResponseModel is 校验用户是否有菜单权限 成功返回结果
type AlibabaCampusAclNewCheckusermenuAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_checkusermenu_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
