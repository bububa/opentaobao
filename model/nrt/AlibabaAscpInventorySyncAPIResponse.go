package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpInventorySyncAPIResponse 库存同步接口 API返回值
// alibaba.ascp.inventory.sync
//
// 新零售联盟商家库存同步接口，用于商家商品库存数量同步到阿里系统
type AlibabaAscpInventorySyncAPIResponse struct {
	model.CommonResponse
	AlibabaAscpInventorySyncAPIResponseModel
}

// AlibabaAscpInventorySyncAPIResponseModel is 库存同步接口 成功返回结果
type AlibabaAscpInventorySyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_inventory_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
