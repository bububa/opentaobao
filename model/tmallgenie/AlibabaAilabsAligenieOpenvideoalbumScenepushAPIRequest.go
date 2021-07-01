package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest
视频专辑场景接入接口 API请求
alibaba.ailabs.aligenie.openvideoalbum.scenepush

视频专辑场景接入接口 */
type AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest struct {
	model.Params
	// 接入场景 0 无应用 1 挂靠应用
	_sceneType int64
	// 如果场景是1 此处为应用id
	_sceneValue string
	// 视频合辑数据
	_paramList []RawVideoAlbum
}

// New
