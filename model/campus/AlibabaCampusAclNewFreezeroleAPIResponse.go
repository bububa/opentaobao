package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewfreezeroleAPIResponse 冻结角色 API返回值
// alibaba.campus.acl.new.freezerole
//
// 冻结角色
type AlibabacampusaclnewfreezeroleAPIResponse struct {
	model.CommonResponse
	AlibabacampusaclnewfreezeroleAPIResponseModel
}

// AlibabacampusaclnewfreezeroleAPIResponseModel is 冻结角色 成功返回结果
type AlibabacampusaclnewfreezeroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_freezerole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
