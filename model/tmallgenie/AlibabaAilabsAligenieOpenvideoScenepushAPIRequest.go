package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieOpenvideoScenepushAPIRequest 视频单集场景接入API API请求
// alibaba.ailabs.aligenie.openvideo.scenepush
//
// 视频单集场景接入API
type AlibabaAilabsAligenieOpenvideoScenepushAPIRequest struct {
	model.Params
	// 待推送的视频数据
	_paramList []RawSingleVideo
	// 挂靠的应用id,在智能应用平台的地址栏可见
	_sceneValue string
	// 内容接入场景0 无应用挂靠 1 应用挂靠
	_sceneType int64
}

// NewAlibabaAilabsAligenieOpenvideoScenepushRequest 初始化AlibabaAilabsAligenieOpenvideoScenepushAPIRequest对象
func NewAlibabaAilabsAligenieOpenvideoScenepushRequest() *AlibabaAilabsAligenieOpenvideoScenepushAPIRequest {
	return &AlibabaAilabsAligenieOpenvideoScenepushAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) Reset() {
	r._paramList = r._paramList[:0]
	r._sceneValue = ""
	r._sceneType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.openvideo.scenepush"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 待推送的视频数据
func (r *AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) SetParamList(_paramList []RawSingleVideo) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) GetParamList() []RawSingleVideo {
	return r._paramList
}

// SetSceneValue is SceneValue Setter
// 挂靠的应用id,在智能应用平台的地址栏可见
func (r *AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) SetSceneValue(_sceneValue string) error {
	r._sceneValue = _sceneValue
	r.Set("scene_value", _sceneValue)
	return nil
}

// GetSceneValue SceneValue Getter
func (r AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) GetSceneValue() string {
	return r._sceneValue
}

// SetSceneType is SceneType Setter
// 内容接入场景0 无应用挂靠 1 应用挂靠
func (r *AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) SetSceneType(_sceneType int64) error {
	r._sceneType = _sceneType
	r.Set("scene_type", _sceneType)
	return nil
}

// GetSceneType SceneType Getter
func (r AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) GetSceneType() int64 {
	return r._sceneType
}

var poolAlibabaAilabsAligenieOpenvideoScenepushAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsAligenieOpenvideoScenepushRequest()
	},
}

// GetAlibabaAilabsAligenieOpenvideoScenepushRequest 从 sync.Pool 获取 AlibabaAilabsAligenieOpenvideoScenepushAPIRequest
func GetAlibabaAilabsAligenieOpenvideoScenepushAPIRequest() *AlibabaAilabsAligenieOpenvideoScenepushAPIRequest {
	return poolAlibabaAilabsAligenieOpenvideoScenepushAPIRequest.Get().(*AlibabaAilabsAligenieOpenvideoScenepushAPIRequest)
}

// ReleaseAlibabaAilabsAligenieOpenvideoScenepushAPIRequest 将 AlibabaAilabsAligenieOpenvideoScenepushAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsAligenieOpenvideoScenepushAPIRequest(v *AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) {
	v.Reset()
	poolAlibabaAilabsAligenieOpenvideoScenepushAPIRequest.Put(v)
}
