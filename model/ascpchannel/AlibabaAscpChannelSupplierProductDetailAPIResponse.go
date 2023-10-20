package ascpchannel

import (
	"encoding/xml"

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

// AlibabaAscpChannelSupplierProductDetailAPIResponseModel is 供应链渠道中心分销品详情查询(供应商专用) 成功返回结果
type AlibabaAscpChannelSupplierProductDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_supplier_product_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAscpChannelSupplierProductDetailResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
