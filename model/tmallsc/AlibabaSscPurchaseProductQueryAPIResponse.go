package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscpurchaseproductqueryAPIResponse 查询已采购的服务产品 API返回值
// alibaba.ssc.purchase.product.query
//
// 查询已采购的服务产品
type AlibabasscpurchaseproductqueryAPIResponse struct {
	model.CommonResponse
	AlibabasscpurchaseproductqueryAPIResponseModel
}

// AlibabasscpurchaseproductqueryAPIResponseModel is 查询已采购的服务产品 成功返回结果
type AlibabasscpurchaseproductqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_purchase_product_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务产品list
	ServiceProducts []ServiceProduct `json:"service_products,omitempty" xml:"service_products>service_product,omitempty"`
	// 错误展示信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
