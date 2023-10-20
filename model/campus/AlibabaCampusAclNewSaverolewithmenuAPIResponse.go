package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewsaverolewithmenuAPIResponse 保存角色级联保存角色和权限的关系 API返回值
// alibaba.campus.acl.new.saverolewithmenu
//
// 保存角色级联保存角色和权限的关系
type AlibabacampusaclnewsaverolewithmenuAPIResponse struct {
	model.CommonResponse
	AlibabacampusaclnewsaverolewithmenuAPIResponseModel
}

// AlibabacampusaclnewsaverolewithmenuAPIResponseModel is 保存角色级联保存角色和权限的关系 成功返回结果
type AlibabacampusaclnewsaverolewithmenuAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_saverolewithmenu_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
