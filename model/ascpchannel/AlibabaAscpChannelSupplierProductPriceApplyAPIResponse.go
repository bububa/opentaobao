package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelSupplierProductPriceApplyAPIResponse 供应商设置渠道产品价格 API返回值
// alibaba.ascp.channel.supplier.product.price.apply
//
// 供应商设置渠道产品价格
type AlibabaAscpChannelSupplierProductPriceApplyAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelSupplierProductPriceApplyAPIResponseModel
}

// AlibabaAscpChannelSupplierProductPriceApplyAPIResponseModel is 供应商设置渠道产品价格 成功返回结果
type AlibabaAscpChannelSupplierProductPriceApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_supplier_product_price_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
