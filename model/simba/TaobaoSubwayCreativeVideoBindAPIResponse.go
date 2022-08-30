package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCreativeVideoBindAPIResponse 绑定视频到创意上 API返回值
// taobao.subway.creative.video.bind
//
// 将用户上传的视频绑定到指定的创意上
type TaobaoSubwayCreativeVideoBindAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayCreativeVideoBindAPIResponseModel
}

// TaobaoSubwayCreativeVideoBindAPIResponseModel is 绑定视频到创意上 成功返回结果
type TaobaoSubwayCreativeVideoBindAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_creative_video_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否绑定成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
