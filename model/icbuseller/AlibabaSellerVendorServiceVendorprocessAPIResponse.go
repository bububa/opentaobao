package icbuseller

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSellerVendorServiceVendorprocessAPIResponse 服务商客户关联信息 API返回值
// alibaba.seller.vendor.service.vendorprocess
//
// 服务商客户关联信息
type AlibabaSellerVendorServiceVendorprocessAPIResponse struct {
	model.CommonResponse
	AlibabaSellerVendorServiceVendorprocessAPIResponseModel
}

// AlibabaSellerVendorServiceVendorprocessAPIResponseModel is 服务商客户关联信息 成功返回结果
type AlibabaSellerVendorServiceVendorprocessAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seller_vendor_service_vendorprocess_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabaSellerVendorServiceVendorprocessResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
