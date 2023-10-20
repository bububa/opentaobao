package icbuseller

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasellervendororderlistAPIResponse 国际站服务市场订单列表接口 API返回值
// alibaba.seller.vendor.order.list
//
// 返回服务商在服务市场的客户订单
type AlibabasellervendororderlistAPIResponse struct {
	model.CommonResponse
	AlibabasellervendororderlistAPIResponseModel
}

// AlibabasellervendororderlistAPIResponseModel is 国际站服务市场订单列表接口 成功返回结果
type AlibabasellervendororderlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seller_vendor_order_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabasellervendororderlistResult `json:"result,omitempty" xml:"result,omitempty"`
}
