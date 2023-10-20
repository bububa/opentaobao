package dutyfree

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadutyfreestockqueryAPIResponse 对外库存查询接口 API返回值
// alibaba.dutyfree.stock.query
//
// 对外部服务提供库存查询接口
type AlibabadutyfreestockqueryAPIResponse struct {
	model.CommonResponse
	AlibabadutyfreestockqueryAPIResponseModel
}

// AlibabadutyfreestockqueryAPIResponseModel is 对外库存查询接口 成功返回结果
type AlibabadutyfreestockqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dutyfree_stock_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabadutyfreestockqueryResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
