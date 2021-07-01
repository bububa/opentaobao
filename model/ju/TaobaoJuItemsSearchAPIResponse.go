package ju

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJuItemsSearchAPIResponse
聚划算商品搜索接口 API返回值
taobao.ju.items.search

搜索聚划算商品 */
type TaobaoJuItemsSearchAPIResponse struct {
	model.CommonResponse
	TaobaoJuItemsSearchAPIResponseModel
}

// TaobaoJuItemsSearchAPIResponseModel is 聚划算商品搜索接口 成功返回结果
type TaobaoJuItemsSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"ju_items_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *PaginationResult `json:"result,omitempty" xml:"result,omitempty"`
}
