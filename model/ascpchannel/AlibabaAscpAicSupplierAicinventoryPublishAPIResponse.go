package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpaicsupplieraicinventorypublishAPIResponse 商家仓操作aic库存发布服务 API返回值
// alibaba.ascp.aic.supplier.aicinventory.publish
//
// 商家调用这个接口来发布增加库存数据
type AlibabaascpaicsupplieraicinventorypublishAPIResponse struct {
	model.CommonResponse
	AlibabaascpaicsupplieraicinventorypublishAPIResponseModel
}

// AlibabaascpaicsupplieraicinventorypublishAPIResponseModel is 商家仓操作aic库存发布服务 成功返回结果
type AlibabaascpaicsupplieraicinventorypublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_aic_supplier_aicinventory_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	AicInventoryPublishResponse *ResultWrapper `json:"aic_inventory_publish_response,omitempty" xml:"aic_inventory_publish_response,omitempty"`
}
