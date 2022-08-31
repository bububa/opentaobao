package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayVideoStateGetAPIResponse 获取视频状态 API返回值
// taobao.subway.video.state.get
//
// 获取已上传视频的状态
type TaobaoSubwayVideoStateGetAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayVideoStateGetAPIResponseModel
}

// TaobaoSubwayVideoStateGetAPIResponseModel is 获取视频状态 成功返回结果
type TaobaoSubwayVideoStateGetAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_video_state_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1, &#34;等待转码&#34;     2, &#34;转码中&#34;     3, &#34;转码失败&#34;     4, &#34;等待审核&#34;     5, &#34;未通过审核&#34;     6, &#34;通过审核&#34;     7, &#34;已删除&#34;     8, &#34;不符合规范&#34;
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}
