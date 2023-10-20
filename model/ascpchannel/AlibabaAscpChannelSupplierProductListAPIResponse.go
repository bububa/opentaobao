package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelSupplierProductListAPIResponse 供应商渠道产品列表查询 API返回值
// alibaba.ascp.channel.supplier.product.list
//
// 供应商查询渠道产品列表
type AlibabaAscpChannelSupplierProductListAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelSupplierProductListAPIResponseModel
}

// AlibabaAscpChannelSupplierProductListAPIResponseModel is 供应商渠道产品列表查询 成功返回结果
type AlibabaAscpChannelSupplierProductListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_supplier_product_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 返回结果
	Module *ProductListQueryResponseForSupplier `json:"module,omitempty" xml:"module,omitempty"`
	// 执行结果是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
