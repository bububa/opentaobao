package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieOpenvideoPushAPIRequest
天猫精灵内容库视频分集数据推送接口 API请求
alibaba.ailabs.aligenie.openvideo.push

天猫精灵内容库视频分集数据推送接口 */
type AlibabaAilabsAligenieOpenvideoPushAPIRequest struct {
	model.Params
	// 待推送的视频数据
	_videos []RawSingleVideo
}

// New
