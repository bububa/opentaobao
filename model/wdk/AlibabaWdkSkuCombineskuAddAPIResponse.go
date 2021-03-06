package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCombineskuAddAPIResponse 组合商品新增接口 API返回值
// alibaba.wdk.sku.combinesku.add
//
// 组合商品新增接口
type AlibabaWdkSkuCombineskuAddAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuCombineskuAddAPIResponseModel
}

// AlibabaWdkSkuCombineskuAddAPIResponseModel is 组合商品新增接口 成功返回结果
type AlibabaWdkSkuCombineskuAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_combinesku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabaWdkSkuCombineskuAddApiResults `json:"result,omitempty" xml:"result,omitempty"`
}
