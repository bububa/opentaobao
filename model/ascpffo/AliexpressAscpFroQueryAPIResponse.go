package ascpffo

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpFroQueryAPIResponse AliExpress销退单查询API API返回值
// aliexpress.ascp.fro.query
//
// AE履约销退单查询接口
type AliexpressAscpFroQueryAPIResponse struct {
	model.CommonResponse
	AliexpressAscpFroQueryAPIResponseModel
}

// AliexpressAscpFroQueryAPIResponseModel is AliExpress销退单查询API 成功返回结果
type AliexpressAscpFroQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_fro_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// DTO
	Result *PageQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
