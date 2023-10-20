package wdkitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemstoreskuqueryAPIResponse 门店商品信息查询 API返回值
// alibaba.wdk.item.storesku.query
//
// 门店商品信息查询
type AlibabawdkitemstoreskuqueryAPIResponse struct {
	model.CommonResponse
	AlibabawdkitemstoreskuqueryAPIResponseModel
}

// AlibabawdkitemstoreskuqueryAPIResponseModel is 门店商品信息查询 成功返回结果
type AlibabawdkitemstoreskuqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_storesku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabawdkitemstoreskuqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
