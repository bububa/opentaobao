package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseQueryloanplansAPIResponse
天猫开新车租后查询还款计划 API返回值
tmall.car.lease.queryloanplans

天猫开新车租后查询还款计划 */
type TmallCarLeaseQueryloanplansAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseQueryloanplansAPIResponseModel
}

// TmallCarLeaseQueryloanplansAPIResponseModel is 天猫开新车租后查询还款计划 成功返回结果
type TmallCarLeaseQueryloanplansAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_queryloanplans_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}
