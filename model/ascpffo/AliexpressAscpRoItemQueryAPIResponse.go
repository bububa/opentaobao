package ascpffo

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascproitemqueryAPIResponse AliExpress退供单明细查询API API返回值
// aliexpress.ascp.ro.item.query
//
// AE仓发 单个退供单明细查询
type AliexpressascproitemqueryAPIResponse struct {
	model.CommonResponse
	AliexpressascproitemqueryAPIResponseModel
}

// AliexpressascproitemqueryAPIResponseModel is AliExpress退供单明细查询API 成功返回结果
type AliexpressascproitemqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_ro_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// dto
	Result *PageQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
