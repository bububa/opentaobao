package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseStatussynchronizeAPIResponse
天猫开新车租后状态同步 API返回值
tmall.car.lease.statussynchronize

天猫开新车租后状态同步 */
type TmallCarLeaseStatussynchronizeAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseStatussynchronizeAPIResponseModel
}

// TmallCarLeaseStatussynchronizeAPIResponseModel is 天猫开新车租后状态同步 成功返回结果
type TmallCarLeaseStatussynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_statussynchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}
