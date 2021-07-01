package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
整车租车回传预约信息 API返回值 
tmall.car.lease.reserve

租赁公司回传预约到店信息
*/
type TmallCarLeaseReserveAPIResponse struct {
    model.CommonResponse
    TmallCarLeaseReserveAPIResponseModel
}

// 整车租车回传预约信息 成功返回结果
type TmallCarLeaseReserveAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_car_lease_reserve_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *TmallCarLeaseReserveResult `json:"result,omitempty" xml:"result,omitempty"`
}
