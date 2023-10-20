package ascpffo

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascproqueryAPIResponse AliExpress退供单查询API API返回值
// aliexpress.ascp.ro.query
//
// AE仓发商家单个退供单查询接口
type AliexpressascproqueryAPIResponse struct {
	model.CommonResponse
	AliexpressascproqueryAPIResponseModel
}

// AliexpressascproqueryAPIResponseModel is AliExpress退供单查询API 成功返回结果
type AliexpressascproqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_ro_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AliexpressascproqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
