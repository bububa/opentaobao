package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieOpenvideoScenepushAPIRequest
视频单集场景接入API API请求
alibaba.ailabs.aligenie.openvideo.scenepush

视频单集场景接入API */
type AlibabaAilabsAligenieOpenvideoScenepushAPIRequest struct {
	model.Params
	// 内容接入场景0 无应用挂靠 1 应用挂靠
	_sceneType int64
	// 挂靠的应用id,在智能应用平台的地址栏可见
	_sceneValue string
	// 待推送的视频数据
	_paramList []RawSingleVideo
}

// New
