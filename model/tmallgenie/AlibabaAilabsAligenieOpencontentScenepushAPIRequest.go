package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsaligenieopencontentscenepushAPIRequest 音频场景接入接口 API请求
// alibaba.ailabs.aligenie.opencontent.scenepush
//
// 天猫精灵音频挂靠场景接入
type AlibabaailabsaligenieopencontentscenepushAPIRequest struct {
	model.Params
	// 如果关联应用此字段为应用id
	_sceneValue string
	// 0 无场景接入  1 关联应用接入
	_sceneType int64
	// 详细内容列表
	_batchContent *BatchContent
}

// NewAlibabaailabsaligenieopencontentscenepushRequest 初始化AlibabaailabsaligenieopencontentscenepushAPIRequest对象
func NewAlibabaailabsaligenieopencontentscenepushRequest() *AlibabaailabsaligenieopencontentscenepushAPIRequest {
	return &AlibabaailabsaligenieopencontentscenepushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsaligenieopencontentscenepushAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.opencontent.scenepush"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsaligenieopencontentscenepushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsaligenieopencontentscenepushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSceneValue is SceneValue Setter
// 如果关联应用此字段为应用id
func (r *AlibabaailabsaligenieopencontentscenepushAPIRequest) SetSceneValue(_sceneValue string) error {
	r._sceneValue = _sceneValue
	r.Set("scene_value", _sceneValue)
	return nil
}

// GetSceneValue SceneValue Getter
func (r AlibabaailabsaligenieopencontentscenepushAPIRequest) GetSceneValue() string {
	return r._sceneValue
}

// SetSceneType is SceneType Setter
// 0 无场景接入  1 关联应用接入
func (r *AlibabaailabsaligenieopencontentscenepushAPIRequest) SetSceneType(_sceneType int64) error {
	r._sceneType = _sceneType
	r.Set("scene_type", _sceneType)
	return nil
}

// GetSceneType SceneType Getter
func (r AlibabaailabsaligenieopencontentscenepushAPIRequest) GetSceneType() int64 {
	return r._sceneType
}

// SetBatchContent is BatchContent Setter
// 详细内容列表
func (r *AlibabaailabsaligenieopencontentscenepushAPIRequest) SetBatchContent(_batchContent *BatchContent) error {
	r._batchContent = _batchContent
	r.Set("batch_content", _batchContent)
	return nil
}

// GetBatchContent BatchContent Getter
func (r AlibabaailabsaligenieopencontentscenepushAPIRequest) GetBatchContent() *BatchContent {
	return r._batchContent
}
