package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthReserveDentalStoresanditemsAPIResponse 查询商户门店，商品列表 API返回值
// alibaba.alihealth.reserve.dental.storesanditems
//
// 查询商户门店，商品列表
type AlibabaAlihealthReserveDentalStoresanditemsAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthReserveDentalStoresanditemsAPIResponseModel
}

// AlibabaAlihealthReserveDentalStoresanditemsAPIResponseModel is 查询商户门店，商品列表 成功返回结果
type AlibabaAlihealthReserveDentalStoresanditemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_reserve_dental_storesanditems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
