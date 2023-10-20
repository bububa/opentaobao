package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsaligenieopenvideoscenepushAPIRequest 视频单集场景接入API API请求
// alibaba.ailabs.aligenie.openvideo.scenepush
//
// 视频单集场景接入API
type AlibabaailabsaligenieopenvideoscenepushAPIRequest struct {
	model.Params
	// 待推送的视频数据
	_paramList []RawSingleVideo
	// 挂靠的应用id,在智能应用平台的地址栏可见
	_sceneValue string
	// 内容接入场景0 无应用挂靠 1 应用挂靠
	_sceneType int64
}

// NewAlibabaailabsaligenieopenvideoscenepushRequest 初始化AlibabaailabsaligenieopenvideoscenepushAPIRequest对象
func NewAlibabaailabsaligenieopenvideoscenepushRequest() *AlibabaailabsaligenieopenvideoscenepushAPIRequest {
	return &AlibabaailabsaligenieopenvideoscenepushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsaligenieopenvideoscenepushAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.openvideo.scenepush"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsaligenieopenvideoscenepushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsaligenieopenvideoscenepushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 待推送的视频数据
func (r *AlibabaailabsaligenieopenvideoscenepushAPIRequest) SetParamList(_paramList []RawSingleVideo) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r AlibabaailabsaligenieopenvideoscenepushAPIRequest) GetParamList() []RawSingleVideo {
	return r._paramList
}

// SetSceneValue is SceneValue Setter
// 挂靠的应用id,在智能应用平台的地址栏可见
func (r *AlibabaailabsaligenieopenvideoscenepushAPIRequest) SetSceneValue(_sceneValue string) error {
	r._sceneValue = _sceneValue
	r.Set("scene_value", _sceneValue)
	return nil
}

// GetSceneValue SceneValue Getter
func (r AlibabaailabsaligenieopenvideoscenepushAPIRequest) GetSceneValue() string {
	return r._sceneValue
}

// SetSceneType is SceneType Setter
// 内容接入场景0 无应用挂靠 1 应用挂靠
func (r *AlibabaailabsaligenieopenvideoscenepushAPIRequest) SetSceneType(_sceneType int64) error {
	r._sceneType = _sceneType
	r.Set("scene_type", _sceneType)
	return nil
}

// GetSceneType SceneType Getter
func (r AlibabaailabsaligenieopenvideoscenepushAPIRequest) GetSceneType() int64 {
	return r._sceneType
}
