package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseConsumeAPIResponse 汽车租赁核销 API返回值
// tmall.car.lease.consume
//
// 租赁公司回传信息，核销
type TmallCarLeaseConsumeAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseConsumeAPIResponseModel
}

// TmallCarLeaseConsumeAPIResponseModel is 汽车租赁核销 成功返回结果
type TmallCarLeaseConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集合
	Result *TmallCarLeaseConsumeResult `json:"result,omitempty" xml:"result,omitempty"`
}
