package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwaycreativevideounbindAPIResponse 创意与视频解绑接口 API返回值
// taobao.subway.creative.video.unbind
//
// 将创意与视频解绑
type TaobaosubwaycreativevideounbindAPIResponse struct {
	model.CommonResponse
	TaobaosubwaycreativevideounbindAPIResponseModel
}

// TaobaosubwaycreativevideounbindAPIResponseModel is 创意与视频解绑接口 成功返回结果
type TaobaosubwaycreativevideounbindAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_creative_video_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否解绑成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
