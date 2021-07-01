package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkStockCabinetQueryAPIResponse
仓内实时库位库存查询 API返回值
alibaba.wdk.stock.cabinet.query

查询仓内实时库位库存信息 */
type AlibabaWdkStockCabinetQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkStockCabinetQueryAPIResponseModel
}

// AlibabaWdkStockCabinetQueryAPIResponseModel is 仓内实时库位库存查询 成功返回结果
type AlibabaWdkStockCabinetQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_stock_cabinet_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkStockCabinetQueryResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
