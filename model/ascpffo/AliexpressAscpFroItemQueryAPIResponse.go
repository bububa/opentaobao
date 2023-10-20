package ascpffo

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascpfroitemqueryAPIResponse AliExpress销退单明细查询API API返回值
// aliexpress.ascp.fro.item.query
//
// AE履约销退单明细查询API
type AliexpressascpfroitemqueryAPIResponse struct {
	model.CommonResponse
	AliexpressascpfroitemqueryAPIResponseModel
}

// AliexpressascpfroitemqueryAPIResponseModel is AliExpress销退单明细查询API 成功返回结果
type AliexpressascpfroitemqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_fro_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AliexpressascpfroitemqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
