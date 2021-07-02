package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthReserveDentalMarkitemAPIResponse 标记商品是否可预约 API返回值
// alibaba.alihealth.reserve.dental.markitem
//
// 标记商品是否可预约
type AlibabaAlihealthReserveDentalMarkitemAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthReserveDentalMarkitemAPIResponseModel
}

// AlibabaAlihealthReserveDentalMarkitemAPIResponseModel is 标记商品是否可预约 成功返回结果
type AlibabaAlihealthReserveDentalMarkitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_reserve_dental_markitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
