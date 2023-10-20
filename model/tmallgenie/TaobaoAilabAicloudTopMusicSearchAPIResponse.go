package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopmusicsearchAPIResponse 对外音乐搜索服务 API返回值
// taobao.ailab.aicloud.top.music.search
//
// 供厂商获取音乐列表
type TaobaoailabaicloudtopmusicsearchAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopmusicsearchAPIResponseModel
}

// TaobaoailabaicloudtopmusicsearchAPIResponseModel is 对外音乐搜索服务 成功返回结果
type TaobaoailabaicloudtopmusicsearchAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_music_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
