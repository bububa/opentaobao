package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkBmPaiyangStockQueryAPIResponse 派样商品门店库存查询接口 API返回值
// alibaba.wdk.bm.paiyang.stock.query
//
// 淘鲜达接入第三方进行派样，第三方查询派样商品的门店库存信息。
type AlibabaWdkBmPaiyangStockQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkBmPaiyangStockQueryAPIResponseModel
}

// AlibabaWdkBmPaiyangStockQueryAPIResponseModel is 派样商品门店库存查询接口 成功返回结果
type AlibabaWdkBmPaiyangStockQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_bm_paiyang_stock_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求出参
	Result *BmResult `json:"result,omitempty" xml:"result,omitempty"`
}
