package wdkitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemStoreskuQueryAPIResponse 门店商品信息查询 API返回值
// alibaba.wdk.item.storesku.query
//
// 门店商品信息查询
type AlibabaWdkItemStoreskuQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemStoreskuQueryAPIResponseModel
}

// AlibabaWdkItemStoreskuQueryAPIResponseModel is 门店商品信息查询 成功返回结果
type AlibabaWdkItemStoreskuQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_storesku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemStoreskuQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
