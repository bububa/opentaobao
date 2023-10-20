package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthReserveDentalModifyrestimeAPIResponse 修改预约时间 API返回值
// alibaba.alihealth.reserve.dental.modifyrestime
//
// 修改预约时间
type AlibabaAlihealthReserveDentalModifyrestimeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthReserveDentalModifyrestimeAPIResponseModel
}

// AlibabaAlihealthReserveDentalModifyrestimeAPIResponseModel is 修改预约时间 成功返回结果
type AlibabaAlihealthReserveDentalModifyrestimeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_reserve_dental_modifyrestime_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihealthReserveDentalModifyrestimeResult `json:"result,omitempty" xml:"result,omitempty"`
}
