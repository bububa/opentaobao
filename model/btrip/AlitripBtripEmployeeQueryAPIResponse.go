package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripemployeequeryAPIResponse 企业员工查询 API返回值
// alitrip.btrip.employee.query
//
// 企业员工查询
type AlitripbtripemployeequeryAPIResponse struct {
	model.CommonResponse
	AlitripbtripemployeequeryAPIResponseModel
}

// AlitripbtripemployeequeryAPIResponseModel is 企业员工查询 成功返回结果
type AlitripbtripemployeequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_employee_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象。
	Result *BtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
