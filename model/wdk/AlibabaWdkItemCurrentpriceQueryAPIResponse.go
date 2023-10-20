package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemcurrentpricequeryAPIResponse 查询商品当前价格 API返回值
// alibaba.wdk.item.currentprice.query
//
// 通过渠道店id/sku编码/渠道查询商品当前价格，一次请求商品数量&lt;=20,返回结果key为skuCode
type AlibabawdkitemcurrentpricequeryAPIResponse struct {
	model.CommonResponse
	AlibabawdkitemcurrentpricequeryAPIResponseModel
}

// AlibabawdkitemcurrentpricequeryAPIResponseModel is 查询商品当前价格 成功返回结果
type AlibabawdkitemcurrentpricequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_currentprice_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabawdkitemcurrentpricequeryResult `json:"result,omitempty" xml:"result,omitempty"`
}
