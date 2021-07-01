package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewDeleteuserroleAPIResponse
删除管理员 API返回值
alibaba.campus.acl.new.deleteuserrole

删除管理员 */
type AlibabaCampusAclNewDeleteuserroleAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewDeleteuserroleAPIResponseModel
}

// AlibabaCampusAclNewDeleteuserroleAPIResponseModel is 删除管理员 成功返回结果
type AlibabaCampusAclNewDeleteuserroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_deleteuserrole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
