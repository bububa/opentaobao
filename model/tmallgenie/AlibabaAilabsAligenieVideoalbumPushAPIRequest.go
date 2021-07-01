package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieVideoalbumPushAPIRequest
天猫精灵内容库视频合辑数据推送接口 API请求
alibaba.ailabs.aligenie.videoalbum.push

三方内容合作厂商可将视频辑数据通过此接口推送至天猫精灵内容库接入中，供天猫精灵使用 */
type AlibabaAilabsAligenieVideoalbumPushAPIRequest struct {
	model.Params
	// 视频合辑数据
	_param1 []RawVideoAlbum
}

// New
