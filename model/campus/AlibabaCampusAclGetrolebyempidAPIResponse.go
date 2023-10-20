package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclgetrolebyempidAPIResponse 根据用户查询角色 API返回值
// alibaba.campus.acl.getrolebyempid
//
// 根据用户查询角色
type AlibabacampusaclgetrolebyempidAPIResponse struct {
	model.CommonResponse
	AlibabacampusaclgetrolebyempidAPIResponseModel
}

// AlibabacampusaclgetrolebyempidAPIResponseModel is 根据用户查询角色 成功返回结果
type AlibabacampusaclgetrolebyempidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_getrolebyempid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CollectionResult `json:"result,omitempty" xml:"result,omitempty"`
}
