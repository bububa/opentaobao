package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
视频单集场景接入API APIRequest
alibaba.ailabs.aligenie.openvideo.scenepush

视频单集场景接入API
*/
type AlibabaAilabsAligenieOpenvideoScenepushRequest struct {
    model.Params

    // 内容接入场景0 无应用挂靠 1 应用挂靠
    sceneType   int64 

    // 挂靠的应用id,在智能应用平台的地址栏可见
    sceneValue   string 

    // 待推送的视频数据
    paramList   []RawSingleVideo 

}

func NewAlibabaAilabsAligenieOpenvideoScenepushRequest() *AlibabaAilabsAligenieOpenvideoScenepushRequest{
    return &AlibabaAilabsAligenieOpenvideoScenepushRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsAligenieOpenvideoScenepushRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.openvideo.scenepush"
}

func (r AlibabaAilabsAligenieOpenvideoScenepushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsAligenieOpenvideoScenepushRequest) SetSceneType(sceneType int64) error {
    r.sceneType = sceneType
    r.Set("scene_type", sceneType)
    return nil
}

func (r AlibabaAilabsAligenieOpenvideoScenepushRequest) GetSceneType() int64 {
    return r.sceneType
}

func (r *AlibabaAilabsAligenieOpenvideoScenepushRequest) SetSceneValue(sceneValue string) error {
    r.sceneValue = sceneValue
    r.Set("scene_value", sceneValue)
    return nil
}

func (r AlibabaAilabsAligenieOpenvideoScenepushRequest) GetSceneValue() string {
    return r.sceneValue
}

func (r *AlibabaAilabsAligenieOpenvideoScenepushRequest) SetParamList(paramList []RawSingleVideo) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

func (r AlibabaAilabsAligenieOpenvideoScenepushRequest) GetParamList() []RawSingleVideo {
    return r.paramList
}

