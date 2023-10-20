package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelSupplierProductDetailAPIResponse 供应链渠道中心分销品详情查询(供应商专用) API返回值
// alibaba.ascp.channel.supplier.product.detail
//
// 供应链渠道中心分销品详情查询(供应商专用)
type AlibabaAscpChannelSupplierProductDetailAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelSupplierProductDetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpChannelSupplierProductDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpChannelSupplierProductDetailAPIResponseModel).Reset()
}

// AlibabaAscpChannelSupplierProductDetailAPIResponseModel is 供应链渠道中心分销品详情查询(供应商专用) 成功返回结果
type AlibabaAscpChannelSupplierProductDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_supplier_product_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAscpChannelSupplierProductDetailResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpChannelSupplierProductDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpChannelSupplierProductDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelSupplierProductDetailAPIResponse)
	},
}

// GetAlibabaAscpChannelSupplierProductDetailAPIResponse 从 sync.Pool 获取 AlibabaAscpChannelSupplierProductDetailAPIResponse
func GetAlibabaAscpChannelSupplierProductDetailAPIResponse() *AlibabaAscpChannelSupplierProductDetailAPIResponse {
	return poolAlibabaAscpChannelSupplierProductDetailAPIResponse.Get().(*AlibabaAscpChannelSupplierProductDetailAPIResponse)
}

// ReleaseAlibabaAscpChannelSupplierProductDetailAPIResponse 将 AlibabaAscpChannelSupplierProductDetailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpChannelSupplierProductDetailAPIResponse(v *AlibabaAscpChannelSupplierProductDetailAPIResponse) {
	v.Reset()
	poolAlibabaAscpChannelSupplierProductDetailAPIResponse.Put(v)
}
