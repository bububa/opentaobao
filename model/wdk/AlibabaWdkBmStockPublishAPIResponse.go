package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkbmstockpublishAPIResponse 品牌营销涉及到的商品的库存同步接口 API返回值
// alibaba.wdk.bm.stock.publish
//
// 用于操作sku的库存
type AlibabawdkbmstockpublishAPIResponse struct {
	model.CommonResponse
	AlibabawdkbmstockpublishAPIResponseModel
}

// AlibabawdkbmstockpublishAPIResponseModel is 品牌营销涉及到的商品的库存同步接口 成功返回结果
type AlibabawdkbmstockpublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_bm_stock_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *BmResult `json:"result,omitempty" xml:"result,omitempty"`
}
