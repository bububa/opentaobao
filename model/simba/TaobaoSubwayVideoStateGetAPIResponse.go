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
	// 1, "等待转码"     2, "转码中"     3, "转码失败"     4, "等待审核"     5, "未通过审核"     6, "通过审核"     7, "已删除"     8, "不符合规范"
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}
