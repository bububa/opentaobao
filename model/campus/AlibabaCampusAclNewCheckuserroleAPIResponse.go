package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewcheckuserroleAPIResponse 校验用户是否有角色 API返回值
// alibaba.campus.acl.new.checkuserrole
//
// 校验用户是否有角色
type AlibabacampusaclnewcheckuserroleAPIResponse struct {
	model.CommonResponse
	AlibabacampusaclnewcheckuserroleAPIResponseModel
}

// AlibabacampusaclnewcheckuserroleAPIResponseModel is 校验用户是否有角色 成功返回结果
type AlibabacampusaclnewcheckuserroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_checkuserrole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
