package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
视频专辑场景接入接口 APIRequest
alibaba.ailabs.aligenie.openvideoalbum.scenepush

视频专辑场景接入接口
*/
type AlibabaAilabsAligenieOpenvideoalbumScenepushRequest struct {
    model.Params

    // 接入场景 0 无应用 1 挂靠应用
    sceneType   int64 

    // 如果场景是1 此处为应用id
    sceneValue   string 

    // 视频合辑数据
    paramList   []RawVideoAlbum 

}

func NewAlibabaAilabsAligenieOpenvideoalbumScenepushRequest() *AlibabaAilabsAligenieOpenvideoalbumScenepushRequest{
    return &AlibabaAilabsAligenieOpenvideoalbumScenepushRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsAligenieOpenvideoalbumScenepushRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.openvideoalbum.scenepush"
}

func (r AlibabaAilabsAligenieOpenvideoalbumScenepushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsAligenieOpenvideoalbumScenepushRequest) SetSceneType(sceneType int64) error {
    r.sceneType = sceneType
    r.Set("scene_type", sceneType)
    return nil
}

func (r AlibabaAilabsAligenieOpenvideoalbumScenepushRequest) GetSceneType() int64 {
    return r.sceneType
}

func (r *AlibabaAilabsAligenieOpenvideoalbumScenepushRequest) SetSceneValue(sceneValue string) error {
    r.sceneValue = sceneValue
    r.Set("scene_value", sceneValue)
    return nil
}

func (r AlibabaAilabsAligenieOpenvideoalbumScenepushRequest) GetSceneValue() string {
    return r.sceneValue
}

func (r *AlibabaAilabsAligenieOpenvideoalbumScenepushRequest) SetParamList(paramList []RawVideoAlbum) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

func (r AlibabaAilabsAligenieOpenvideoalbumScenepushRequest) GetParamList() []RawVideoAlbum {
    return r.paramList
}

