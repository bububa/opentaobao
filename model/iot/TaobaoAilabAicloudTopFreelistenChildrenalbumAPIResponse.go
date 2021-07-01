package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse
儿童音频列表 API返回值
taobao.ailab.aicloud.top.freelisten.childrenalbum

儿童音频列表 */
type TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponseModel
}

// TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponseModel is 儿童音频列表 成功返回结果
type TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_freelisten_childrenalbum_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
