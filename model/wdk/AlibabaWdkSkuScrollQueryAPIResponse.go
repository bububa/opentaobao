package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuScrollQueryAPIResponse 门店商品批量游标方式查询接口 API返回值
// alibaba.wdk.sku.scroll.query
//
// 通过游标方式批量获取门店商品信息，包括商品条码，商品名称，价格，会员价等信息。
type AlibabaWdkSkuScrollQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuScrollQueryAPIResponseModel
}

// AlibabaWdkSkuScrollQueryAPIResponseModel is 门店商品批量游标方式查询接口 成功返回结果
type AlibabaWdkSkuScrollQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_scroll_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ApiScrollPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
