package ascpffo

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascpwarehouseinventoryqueryAPIResponse AliExpress在仓库存查询API API返回值
// aliexpress.ascp.warehouse.inventory.query
//
// AliExpress在仓库存查询查询API
type AliexpressascpwarehouseinventoryqueryAPIResponse struct {
	model.CommonResponse
	AliexpressascpwarehouseinventoryqueryAPIResponseModel
}

// AliexpressascpwarehouseinventoryqueryAPIResponseModel is AliExpress在仓库存查询API 成功返回结果
type AliexpressascpwarehouseinventoryqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_warehouse_inventory_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页查询结果
	Result *PageQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
