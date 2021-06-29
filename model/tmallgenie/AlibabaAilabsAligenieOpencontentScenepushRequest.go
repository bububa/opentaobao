package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
音频场景接入接口 API请求
alibaba.ailabs.aligenie.opencontent.scenepush

天猫精灵音频挂靠场景接入
*/
type AlibabaAilabsAligenieOpencontentScenepushRequest struct {
    model.Params
    // 0 无场景接入  1 关联应用接入
    sceneType   int64
    // 如果关联应用此字段为应用id
    sceneValue   string
    // 详细内容列表
    batchContent   *BatchContent
}

// 初始化AlibabaAilabsAligenieOpencontentScenepushRequest对象
func NewAlibabaAilabsAligenieOpencontentScenepushRequest() *AlibabaAilabsAligenieOpencontentScenepushRequest{
    return &AlibabaAilabsAligenieOpencontentScenepushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieOpencontentScenepushRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.opencontent.scenepush"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieOpencontentScenepushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SceneType Setter
// 0 无场景接入  1 关联应用接入
func (r *AlibabaAilabsAligenieOpencontentScenepushRequest) SetSceneType(sceneType int64) error {
    r.sceneType = sceneType
    r.Set("scene_type", sceneType)
    return nil
}

// SceneType Getter
func (r AlibabaAilabsAligenieOpencontentScenepushRequest) GetSceneType() int64 {
    return r.sceneType
}
// SceneValue Setter
// 如果关联应用此字段为应用id
func (r *AlibabaAilabsAligenieOpencontentScenepushRequest) SetSceneValue(sceneValue string) error {
    r.sceneValue = sceneValue
    r.Set("scene_value", sceneValue)
    return nil
}

// SceneValue Getter
func (r AlibabaAilabsAligenieOpencontentScenepushRequest) GetSceneValue() string {
    return r.sceneValue
}
// BatchContent Setter
// 详细内容列表
func (r *AlibabaAilabsAligenieOpencontentScenepushRequest) SetBatchContent(batchContent *BatchContent) error {
    r.batchContent = batchContent
    r.Set("batch_content", batchContent)
    return nil
}

// BatchContent Getter
func (r AlibabaAilabsAligenieOpencontentScenepushRequest) GetBatchContent() *BatchContent {
    return r.batchContent
}
