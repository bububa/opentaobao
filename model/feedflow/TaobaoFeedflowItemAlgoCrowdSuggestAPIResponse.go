package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAlgoCrowdSuggestAPIResponse 单品人群建议出价 API返回值
// taobao.feedflow.item.algo.crowd.suggest
//
// 给超级推荐的广告主查看建议出价
type TaobaoFeedflowItemAlgoCrowdSuggestAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemAlgoCrowdSuggestAPIResponseModel
}

// TaobaoFeedflowItemAlgoCrowdSuggestAPIResponseModel is 单品人群建议出价 成功返回结果
type TaobaoFeedflowItemAlgoCrowdSuggestAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_algo_crowd_suggest_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemAlgoCrowdSuggestResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
