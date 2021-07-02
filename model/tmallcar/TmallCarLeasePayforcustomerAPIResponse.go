package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeasePayforcustomerAPIResponse 天猫开新车租后代客户还款 API返回值
// tmall.car.lease.payforcustomer
//
// 天猫开新车租后代客户还款
type TmallCarLeasePayforcustomerAPIResponse struct {
	model.CommonResponse
	TmallCarLeasePayforcustomerAPIResponseModel
}

// TmallCarLeasePayforcustomerAPIResponseModel is 天猫开新车租后代客户还款 成功返回结果
type TmallCarLeasePayforcustomerAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_payforcustomer_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}
