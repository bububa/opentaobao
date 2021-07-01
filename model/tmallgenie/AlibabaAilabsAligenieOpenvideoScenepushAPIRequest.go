package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
视频单集场景接入API API请求
alibaba.ailabs.aligenie.openvideo.scenepush

视频单集场景接入API
*/
type AlibabaAilabsAligenieOpenvideoScenepushAPIRequest struct {
    model.Params
    // 内容接入场景0 无应用挂靠 1 应用挂靠
    _sceneType   int64
    // 挂靠的应用id,在智能应用平台的地址栏可见
    _sceneValue   string
    // 待推送的视频数据
    _paramList   []RawSingleVideo
}

// 初始化AlibabaAilabsAligenieOpenvideoScenepushAPIRequest对象
func NewAlibabaAilabsAligenieOpenvideoScenepushRequest() *AlibabaAilabsAligenieOpenvideoScenepushAPIRequest{
    return &AlibabaAilabsAligenieOpenvideoScenepushAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.openvideo.scenepush"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SceneType Setter
// 内容接入场景0 无应用挂靠 1 应用挂靠
func (r *AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) SetSceneType(_sceneType int64) error {
    r._sceneType = _sceneType
    r.Set("scene_type", _sceneType)
    return nil
}

// SceneType Getter
func (r AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) GetSceneType() int64 {
    return r._sceneType
}
// SceneValue Setter
// 挂靠的应用id,在智能应用平台的地址栏可见
func (r *AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) SetSceneValue(_sceneValue string) error {
    r._sceneValue = _sceneValue
    r.Set("scene_value", _sceneValue)
    return nil
}

// SceneValue Getter
func (r AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) GetSceneValue() string {
    return r._sceneValue
}
// ParamList Setter
// 待推送的视频数据
func (r *AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) SetParamList(_paramList []RawSingleVideo) error {
    r._paramList = _paramList
    r.Set("param_list", _paramList)
    return nil
}

// ParamList Getter
func (r AlibabaAilabsAligenieOpenvideoScenepushAPIRequest) GetParamList() []RawSingleVideo {
    return r._paramList
}
