package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest 视频专辑场景接入接口 API请求
// alibaba.ailabs.aligenie.openvideoalbum.scenepush
//
// 视频专辑场景接入接口
type AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest struct {
	model.Params
	// 视频合辑数据
	_paramList []RawVideoAlbum
	// 如果场景是1 此处为应用id
	_sceneValue string
	// 接入场景 0 无应用 1 挂靠应用
	_sceneType int64
}

// NewAlibabaAilabsAligenieOpenvideoalbumScenepushRequest 初始化AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest对象
func NewAlibabaAilabsAligenieOpenvideoalbumScenepushRequest() *AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest {
	return &AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.openvideoalbum.scenepush"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 视频合辑数据
func (r *AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest) SetParamList(_paramList []RawVideoAlbum) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest) GetParamList() []RawVideoAlbum {
	return r._paramList
}

// SetSceneValue is SceneValue Setter
// 如果场景是1 此处为应用id
func (r *AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest) SetSceneValue(_sceneValue string) error {
	r._sceneValue = _sceneValue
	r.Set("scene_value", _sceneValue)
	return nil
}

// GetSceneValue SceneValue Getter
func (r AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest) GetSceneValue() string {
	return r._sceneValue
}

// SetSceneType is SceneType Setter
// 接入场景 0 无应用 1 挂靠应用
func (r *AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest) SetSceneType(_sceneType int64) error {
	r._sceneType = _sceneType
	r.Set("scene_type", _sceneType)
	return nil
}

// GetSceneType SceneType Getter
func (r AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest) GetSceneType() int64 {
	return r._sceneType
}
