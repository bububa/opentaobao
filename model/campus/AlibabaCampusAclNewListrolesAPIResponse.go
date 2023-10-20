package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewlistrolesAPIResponse 查询全部角色 API返回值
// alibaba.campus.acl.new.listroles
//
// 查询全部角色
type AlibabacampusaclnewlistrolesAPIResponse struct {
	model.CommonResponse
	AlibabacampusaclnewlistrolesAPIResponseModel
}

// AlibabacampusaclnewlistrolesAPIResponseModel is 查询全部角色 成功返回结果
type AlibabacampusaclnewlistrolesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_listroles_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}
