package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieOpencontentScenepushAPIRequest
音频场景接入接口 API请求
alibaba.ailabs.aligenie.opencontent.scenepush

天猫精灵音频挂靠场景接入 */
type AlibabaAilabsAligenieOpencontentScenepushAPIRequest struct {
	model.Params
	// 0 无场景接入  1 关联应用接入
	_sceneType int64
	// 如果关联应用此字段为应用id
	_sceneValue string
	// 详细内容列表
	_batchContent *BatchContent
}

// New
