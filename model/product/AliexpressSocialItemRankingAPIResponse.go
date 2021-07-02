package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSocialItemRankingAPIResponse 社交排行榜 API返回值
// aliexpress.social.item.ranking
//
// 社交商品成交排行榜
type AliexpressSocialItemRankingAPIResponse struct {
	model.CommonResponse
	AliexpressSocialItemRankingAPIResponseModel
}

// AliexpressSocialItemRankingAPIResponseModel is 社交排行榜 成功返回结果
type AliexpressSocialItemRankingAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_social_item_ranking_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类型
	Result *ItemPickPagingResult `json:"result,omitempty" xml:"result,omitempty"`
}
