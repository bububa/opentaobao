package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpaicsupplieraicinventorynegativesalepublishAPIResponse AIC负卖库存新增和修改接口 API返回值
// alibaba.ascp.aic.supplier.aicinventory.negative.sale.publish
//
// 新增负卖库存记录和变更负卖库存记录
type AlibabaascpaicsupplieraicinventorynegativesalepublishAPIResponse struct {
	model.CommonResponse
	AlibabaascpaicsupplieraicinventorynegativesalepublishAPIResponseModel
}

// AlibabaascpaicsupplieraicinventorynegativesalepublishAPIResponseModel is AIC负卖库存新增和修改接口 成功返回结果
type AlibabaascpaicsupplieraicinventorynegativesalepublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_aic_supplier_aicinventory_negative_sale_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	FutureInvItemResponse *ResultWrapper `json:"future_inv_item_response,omitempty" xml:"future_inv_item_response,omitempty"`
}
