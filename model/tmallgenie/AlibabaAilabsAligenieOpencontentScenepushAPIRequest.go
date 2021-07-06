package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieOpencontentScenepushAPIRequest 音频场景接入接口 API请求
// alibaba.ailabs.aligenie.opencontent.scenepush
//
// 天猫精灵音频挂靠场景接入
type AlibabaAilabsAligenieOpencontentScenepushAPIRequest struct {
	model.Params
	// 如果关联应用此字段为应用id
	_sceneValue string
	// 0 无场景接入  1 关联应用接入
	_sceneType int64
	// 详细内容列表
	_batchContent *BatchContent
}

// NewAlibabaAilabsAligenieOpencontentScenepushRequest 初始化AlibabaAilabsAligenieOpencontentScenepushAPIRequest对象
func NewAlibabaAilabsAligenieOpencontentScenepushRequest() *AlibabaAilabsAligenieOpencontentScenepushAPIRequest {
	return &AlibabaAilabsAligenieOpencontentScenepushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieOpencontentScenepushAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.opencontent.scenepush"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieOpencontentScenepushAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSceneValue is SceneValue Setter
// 如果关联应用此字段为应用id
func (r *AlibabaAilabsAligenieOpencontentScenepushAPIRequest) SetSceneValue(_sceneValue string) error {
	r._sceneValue = _sceneValue
	r.Set("scene_value", _sceneValue)
	return nil
}

// GetSceneValue SceneValue Getter
func (r AlibabaAilabsAligenieOpencontentScenepushAPIRequest) GetSceneValue() string {
	return r._sceneValue
}

// SetSceneType is SceneType Setter
// 0 无场景接入  1 关联应用接入
func (r *AlibabaAilabsAligenieOpencontentScenepushAPIRequest) SetSceneType(_sceneType int64) error {
	r._sceneType = _sceneType
	r.Set("scene_type", _sceneType)
	return nil
}

// GetSceneType SceneType Getter
func (r AlibabaAilabsAligenieOpencontentScenepushAPIRequest) GetSceneType() int64 {
	return r._sceneType
}

// SetBatchContent is BatchContent Setter
// 详细内容列表
func (r *AlibabaAilabsAligenieOpencontentScenepushAPIRequest) SetBatchContent(_batchContent *BatchContent) error {
	r._batchContent = _batchContent
	r.Set("batch_content", _batchContent)
	return nil
}

// GetBatchContent BatchContent Getter
func (r AlibabaAilabsAligenieOpencontentScenepushAPIRequest) GetBatchContent() *BatchContent {
	return r._batchContent
}
