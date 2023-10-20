package ascpffo

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascpinventorylogqueryAPIResponse AliExpress库存流水查询API API返回值
// aliexpress.ascp.inventory.log.query
//
// AliExpress库存流水查询API
type AliexpressascpinventorylogqueryAPIResponse struct {
	model.CommonResponse
	AliexpressascpinventorylogqueryAPIResponseModel
}

// AliexpressascpinventorylogqueryAPIResponseModel is AliExpress库存流水查询API 成功返回结果
type AliexpressascpinventorylogqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_inventory_log_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询出参
	Result *PageQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
