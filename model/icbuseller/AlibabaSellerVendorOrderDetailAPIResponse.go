package icbuseller

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasellervendororderdetailAPIResponse 国际站服务市场订单详情接口 API返回值
// alibaba.seller.vendor.order.detail
//
// 国际站服务市场订单列表接口
type AlibabasellervendororderdetailAPIResponse struct {
	model.CommonResponse
	AlibabasellervendororderdetailAPIResponseModel
}

// AlibabasellervendororderdetailAPIResponseModel is 国际站服务市场订单详情接口 成功返回结果
type AlibabasellervendororderdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seller_vendor_order_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回对象
	Result *AlibabasellervendororderdetailResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
