package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwayitemvideouploadAPIResponse 创意视频上传 API返回值
// taobao.subway.item.video.upload
//
// 为用户提供视频上传的功能
type TaobaosubwayitemvideouploadAPIResponse struct {
	model.CommonResponse
	TaobaosubwayitemvideouploadAPIResponseModel
}

// TaobaosubwayitemvideouploadAPIResponseModel is 创意视频上传 成功返回结果
type TaobaosubwayitemvideouploadAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_item_video_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 视频信息
	Result *VideoFeedDto `json:"result,omitempty" xml:"result,omitempty"`
}
