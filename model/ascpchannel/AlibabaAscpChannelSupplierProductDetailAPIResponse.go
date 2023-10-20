package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelsupplierproductdetailAPIResponse 供应链渠道中心分销品详情查询(供应商专用) API返回值
// alibaba.ascp.channel.supplier.product.detail
//
// 供应链渠道中心分销品详情查询(供应商专用)
type AlibabaascpchannelsupplierproductdetailAPIResponse struct {
	model.CommonResponse
	AlibabaascpchannelsupplierproductdetailAPIResponseModel
}

// AlibabaascpchannelsupplierproductdetailAPIResponseModel is 供应链渠道中心分销品详情查询(供应商专用) 成功返回结果
type AlibabaascpchannelsupplierproductdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_supplier_product_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaascpchannelsupplierproductdetailResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
