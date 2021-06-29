package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
视频专辑场景接入接口 API请求
alibaba.ailabs.aligenie.openvideoalbum.scenepush

视频专辑场景接入接口
*/
type AlibabaAilabsAligenieOpenvideoalbumScenepushRequest struct {
    model.Params
    // 接入场景 0 无应用 1 挂靠应用
    _sceneType   int64
    // 如果场景是1 此处为应用id
    _sceneValue   string
    // 视频合辑数据
    _paramList   []RawVideoAlbum
}

// 初始化AlibabaAilabsAligenieOpenvideoalbumScenepushRequest对象
func NewAlibabaAilabsAligenieOpenvideoalbumScenepushRequest() *AlibabaAilabsAligenieOpenvideoalbumScenepushRequest{
    return &AlibabaAilabsAligenieOpenvideoalbumScenepushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieOpenvideoalbumScenepushRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.openvideoalbum.scenepush"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieOpenvideoalbumScenepushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SceneType Setter
// 接入场景 0 无应用 1 挂靠应用
func (r *AlibabaAilabsAligenieOpenvideoalbumScenepushRequest) SetSceneType(_sceneType int64) error {
    r._sceneType = _sceneType
    r.Set("scene_type", _sceneType)
    return nil
}

// SceneType Getter
func (r AlibabaAilabsAligenieOpenvideoalbumScenepushRequest) GetSceneType() int64 {
    return r._sceneType
}
// SceneValue Setter
// 如果场景是1 此处为应用id
func (r *AlibabaAilabsAligenieOpenvideoalbumScenepushRequest) SetSceneValue(_sceneValue string) error {
    r._sceneValue = _sceneValue
    r.Set("scene_value", _sceneValue)
    return nil
}

// SceneValue Getter
func (r AlibabaAilabsAligenieOpenvideoalbumScenepushRequest) GetSceneValue() string {
    return r._sceneValue
}
// ParamList Setter
// 视频合辑数据
func (r *AlibabaAilabsAligenieOpenvideoalbumScenepushRequest) SetParamList(_paramList []RawVideoAlbum) error {
    r._paramList = _paramList
    r.Set("param_list", _paramList)
    return nil
}

// ParamList Getter
func (r AlibabaAilabsAligenieOpenvideoalbumScenepushRequest) GetParamList() []RawVideoAlbum {
    return r._paramList
}
