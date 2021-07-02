package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSocialItemSearchAPIResponse AE社交选品 API返回值
// aliexpress.social.item.search
//
// AE社交选品,通过各种筛选条件对社交商品池进行筛选
type AliexpressSocialItemSearchAPIResponse struct {
	model.CommonResponse
	AliexpressSocialItemSearchAPIResponseModel
}

// AliexpressSocialItemSearchAPIResponseModel is AE社交选品 成功返回结果
type AliexpressSocialItemSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_social_item_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 报类型
	Result *ItemPickPagingResult `json:"result,omitempty" xml:"result,omitempty"`
}
