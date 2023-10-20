package icbuseller

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSellerVendorTradePurchaseAPIResponse 查看购买人的订单记录以及授权时间 API返回值
// alibaba.seller.vendor.trade.purchase
//
// 查看购买人的订单记录以及授权时间
type AlibabaSellerVendorTradePurchaseAPIResponse struct {
	model.CommonResponse
	AlibabaSellerVendorTradePurchaseAPIResponseModel
}

// AlibabaSellerVendorTradePurchaseAPIResponseModel is 查看购买人的订单记录以及授权时间 成功返回结果
type AlibabaSellerVendorTradePurchaseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seller_vendor_trade_purchase_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabaSellerVendorTradePurchaseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
