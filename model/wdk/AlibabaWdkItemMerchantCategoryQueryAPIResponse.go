package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemmerchantcategoryqueryAPIResponse 查询商品的商家叶子类目 API返回值
// alibaba.wdk.item.merchant.category.query
//
// 查询商品的商家叶子类目
type AlibabawdkitemmerchantcategoryqueryAPIResponse struct {
	model.CommonResponse
	AlibabawdkitemmerchantcategoryqueryAPIResponseModel
}

// AlibabawdkitemmerchantcategoryqueryAPIResponseModel is 查询商品的商家叶子类目 成功返回结果
type AlibabawdkitemmerchantcategoryqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_merchant_category_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *WdkOpenSkuMerchantCatServiceQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
