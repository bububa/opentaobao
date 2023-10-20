package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsaligenieopenvideoalbumscenepushAPIRequest 视频专辑场景接入接口 API请求
// alibaba.ailabs.aligenie.openvideoalbum.scenepush
//
// 视频专辑场景接入接口
type AlibabaailabsaligenieopenvideoalbumscenepushAPIRequest struct {
	model.Params
	// 视频合辑数据
	_paramList []RawVideoAlbum
	// 如果场景是1 此处为应用id
	_sceneValue string
	// 接入场景 0 无应用 1 挂靠应用
	_sceneType int64
}

// NewAlibabaailabsaligenieopenvideoalbumscenepushRequest 初始化AlibabaailabsaligenieopenvideoalbumscenepushAPIRequest对象
func NewAlibabaailabsaligenieopenvideoalbumscenepushRequest() *AlibabaailabsaligenieopenvideoalbumscenepushAPIRequest {
	return &AlibabaailabsaligenieopenvideoalbumscenepushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsaligenieopenvideoalbumscenepushAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.openvideoalbum.scenepush"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsaligenieopenvideoalbumscenepushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsaligenieopenvideoalbumscenepushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 视频合辑数据
func (r *AlibabaailabsaligenieopenvideoalbumscenepushAPIRequest) SetParamList(_paramList []RawVideoAlbum) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r AlibabaailabsaligenieopenvideoalbumscenepushAPIRequest) GetParamList() []RawVideoAlbum {
	return r._paramList
}

// SetSceneValue is SceneValue Setter
// 如果场景是1 此处为应用id
func (r *AlibabaailabsaligenieopenvideoalbumscenepushAPIRequest) SetSceneValue(_sceneValue string) error {
	r._sceneValue = _sceneValue
	r.Set("scene_value", _sceneValue)
	return nil
}

// GetSceneValue SceneValue Getter
func (r AlibabaailabsaligenieopenvideoalbumscenepushAPIRequest) GetSceneValue() string {
	return r._sceneValue
}

// SetSceneType is SceneType Setter
// 接入场景 0 无应用 1 挂靠应用
func (r *AlibabaailabsaligenieopenvideoalbumscenepushAPIRequest) SetSceneType(_sceneType int64) error {
	r._sceneType = _sceneType
	r.Set("scene_type", _sceneType)
	return nil
}

// GetSceneType SceneType Getter
func (r AlibabaailabsaligenieopenvideoalbumscenepushAPIRequest) GetSceneType() int64 {
	return r._sceneType
}
