package icbulogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaonetouchlogisticsexpressorderdetailgetAPIResponse 订单详细信息(面单及仓库信息) API返回值
// alibaba.onetouch.logistics.express.order.detail.get
//
// 订单详细信息(面单及仓库信息)
type AlibabaonetouchlogisticsexpressorderdetailgetAPIResponse struct {
	model.CommonResponse
	AlibabaonetouchlogisticsexpressorderdetailgetAPIResponseModel
}

// AlibabaonetouchlogisticsexpressorderdetailgetAPIResponseModel is 订单详细信息(面单及仓库信息) 成功返回结果
type AlibabaonetouchlogisticsexpressorderdetailgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_order_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaonetouchlogisticsexpressorderdetailgetResult `json:"result,omitempty" xml:"result,omitempty"`
}
