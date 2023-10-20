package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelSupplierProductGoodsBindAPIResponse 渠道产品与货品绑定接口 API返回值
// alibaba.ascp.channel.supplier.product.goods.bind
//
// 渠道产品与货品绑定接口
type AlibabaAscpChannelSupplierProductGoodsBindAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelSupplierProductGoodsBindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpChannelSupplierProductGoodsBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpChannelSupplierProductGoodsBindAPIResponseModel).Reset()
}

// AlibabaAscpChannelSupplierProductGoodsBindAPIResponseModel is 渠道产品与货品绑定接口 成功返回结果
type AlibabaAscpChannelSupplierProductGoodsBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_supplier_product_goods_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpChannelSupplierProductGoodsBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpChannelSupplierProductGoodsBindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelSupplierProductGoodsBindAPIResponse)
	},
}

// GetAlibabaAscpChannelSupplierProductGoodsBindAPIResponse 从 sync.Pool 获取 AlibabaAscpChannelSupplierProductGoodsBindAPIResponse
func GetAlibabaAscpChannelSupplierProductGoodsBindAPIResponse() *AlibabaAscpChannelSupplierProductGoodsBindAPIResponse {
	return poolAlibabaAscpChannelSupplierProductGoodsBindAPIResponse.Get().(*AlibabaAscpChannelSupplierProductGoodsBindAPIResponse)
}

// ReleaseAlibabaAscpChannelSupplierProductGoodsBindAPIResponse 将 AlibabaAscpChannelSupplierProductGoodsBindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpChannelSupplierProductGoodsBindAPIResponse(v *AlibabaAscpChannelSupplierProductGoodsBindAPIResponse) {
	v.Reset()
	poolAlibabaAscpChannelSupplierProductGoodsBindAPIResponse.Put(v)
}
