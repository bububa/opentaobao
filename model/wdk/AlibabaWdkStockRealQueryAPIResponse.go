package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkStockRealQueryAPIResponse 仓内实时库存查询 API返回值
// alibaba.wdk.stock.real.query
//
// 查询仓内实时库存信息
type AlibabaWdkStockRealQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkStockRealQueryAPIResponseModel
}

// AlibabaWdkStockRealQueryAPIResponseModel is 仓内实时库存查询 成功返回结果
type AlibabaWdkStockRealQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_stock_real_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaWdkStockRealQueryResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
