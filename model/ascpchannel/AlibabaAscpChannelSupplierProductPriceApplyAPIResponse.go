package ascpchannel

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAscpChannelSupplierProductPriceApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpChannelSupplierProductPriceApplyAPIResponseModel).Reset()
}

// AlibabaAscpChannelSupplierProductPriceApplyAPIResponseModel is 供应商设置渠道产品价格 成功返回结果
type AlibabaAscpChannelSupplierProductPriceApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_supplier_product_price_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpChannelSupplierProductPriceApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpChannelSupplierProductPriceApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelSupplierProductPriceApplyAPIResponse)
	},
}

// GetAlibabaAscpChannelSupplierProductPriceApplyAPIResponse 从 sync.Pool 获取 AlibabaAscpChannelSupplierProductPriceApplyAPIResponse
func GetAlibabaAscpChannelSupplierProductPriceApplyAPIResponse() *AlibabaAscpChannelSupplierProductPriceApplyAPIResponse {
	return poolAlibabaAscpChannelSupplierProductPriceApplyAPIResponse.Get().(*AlibabaAscpChannelSupplierProductPriceApplyAPIResponse)
}

// ReleaseAlibabaAscpChannelSupplierProductPriceApplyAPIResponse 将 AlibabaAscpChannelSupplierProductPriceApplyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpChannelSupplierProductPriceApplyAPIResponse(v *AlibabaAscpChannelSupplierProductPriceApplyAPIResponse) {
	v.Reset()
	poolAlibabaAscpChannelSupplierProductPriceApplyAPIResponse.Put(v)
}
