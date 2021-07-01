package alihealthlab

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthReservationOrderVerifyAPIResponse
预约单核销接口 API返回值
alibaba.alihealth.reservation.order.verify

预约单核销 */
type AlibabaAlihealthReservationOrderVerifyAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthReservationOrderVerifyAPIResponseModel
}

// AlibabaAlihealthReservationOrderVerifyAPIResponseModel is 预约单核销接口 成功返回结果
type AlibabaAlihealthReservationOrderVerifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_reservation_order_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// SUCCESS - 成功，FAIL - 失败，UNKNOWN - 未知或处理中
	ResultStatus string `json:"result_status,omitempty" xml:"result_status,omitempty"`
	// 00000为成功 其它错误码 https://www.yuque.com/alibabahealth/mbhk06/gnvi8b#ICrXp
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}
