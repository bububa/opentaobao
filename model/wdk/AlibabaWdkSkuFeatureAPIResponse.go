package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuFeatureAPIResponse 商品标记接口 API返回值
// alibaba.wdk.sku.feature
//
// 给淘鲜达商品属性之外的打标通用能力，满足商品一些特殊的需求，比如是否参加营销。
type AlibabaWdkSkuFeatureAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuFeatureAPIResponseModel
}

// AlibabaWdkSkuFeatureAPIResponseModel is 商品标记接口 成功返回结果
type AlibabaWdkSkuFeatureAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_feature_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	Result *AlibabaWdkSkuFeatureApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
