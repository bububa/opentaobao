package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopLikeListAPIResponse
列出收藏列表 API返回值
taobao.ailab.aicloud.top.like.list

列出收藏列表 */
type TaobaoAilabAicloudTopLikeListAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopLikeListAPIResponseModel
}

// TaobaoAilabAicloudTopLikeListAPIResponseModel is 列出收藏列表 成功返回结果
type TaobaoAilabAicloudTopLikeListAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_like_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
