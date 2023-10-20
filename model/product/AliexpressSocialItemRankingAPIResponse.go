package product

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AliexpressSocialItemRankingAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSocialItemRankingAPIResponseModel).Reset()
}

// AliexpressSocialItemRankingAPIResponseModel is 社交排行榜 成功返回结果
type AliexpressSocialItemRankingAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_social_item_ranking_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类型
	Result *ItemPickPagingResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSocialItemRankingAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressSocialItemRankingAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSocialItemRankingAPIResponse)
	},
}

// GetAliexpressSocialItemRankingAPIResponse 从 sync.Pool 获取 AliexpressSocialItemRankingAPIResponse
func GetAliexpressSocialItemRankingAPIResponse() *AliexpressSocialItemRankingAPIResponse {
	return poolAliexpressSocialItemRankingAPIResponse.Get().(*AliexpressSocialItemRankingAPIResponse)
}

// ReleaseAliexpressSocialItemRankingAPIResponse 将 AliexpressSocialItemRankingAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSocialItemRankingAPIResponse(v *AliexpressSocialItemRankingAPIResponse) {
	v.Reset()
	poolAliexpressSocialItemRankingAPIResponse.Put(v)
}
