package aeusergrowth

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressusergrowthsearchitemsgetAPIResponse 第三方平台搜索AE商品 API返回值
// aliexpress.usergrowth.search.items.get
//
// 第三方平台的搜索服务   获取AE商品list
type AliexpressusergrowthsearchitemsgetAPIResponse struct {
	model.CommonResponse
	AliexpressusergrowthsearchitemsgetAPIResponseModel
}

// AliexpressusergrowthsearchitemsgetAPIResponseModel is 第三方平台搜索AE商品 成功返回结果
type AliexpressusergrowthsearchitemsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_usergrowth_search_items_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// response model
	Result *AliexpressusergrowthsearchitemsgetResult `json:"result,omitempty" xml:"result,omitempty"`
}
