package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaToolsItemsTopGetAPIResponse
取得一个关键词的推广组排名列表 API返回值
taobao.simba.tools.items.top.get

取得一个关键词的推广组排名列表 */
type TaobaoSimbaToolsItemsTopGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaToolsItemsTopGetAPIResponseModel
}

// TaobaoSimbaToolsItemsTopGetAPIResponseModel is 取得一个关键词的推广组排名列表 成功返回结果
type TaobaoSimbaToolsItemsTopGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_tools_items_top_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广组信息列表
	Rankeditems []RankedItem `json:"rankeditems,omitempty" xml:"rankeditems>ranked_item,omitempty"`
}
