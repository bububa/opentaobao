package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵内容接入标准接口 APIRequest
alibaba.ailabs.aligenie.opencontent.push

第三方内容接入天猫精灵内容库，供相关技能使用
*/
type AlibabaAilabsAligenieOpencontentPushRequest struct {
    model.Params

    // 在Aligenie开放平台创建的技能的ID
    skillId   int64 

    // 详细内容列表
    contents   *BatchContent 

}

func NewAlibabaAilabsAligenieOpencontentPushRequest() *AlibabaAilabsAligenieOpencontentPushRequest{
    return &AlibabaAilabsAligenieOpencontentPushRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsAligenieOpencontentPushRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.opencontent.push"
}

func (r AlibabaAilabsAligenieOpencontentPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsAligenieOpencontentPushRequest) SetSkillId(skillId int64) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

func (r AlibabaAilabsAligenieOpencontentPushRequest) GetSkillId() int64 {
    return r.skillId
}

func (r *AlibabaAilabsAligenieOpencontentPushRequest) SetContents(contents *BatchContent) error {
    r.contents = contents
    r.Set("contents", contents)
    return nil
}

func (r AlibabaAilabsAligenieOpencontentPushRequest) GetContents() *BatchContent {
    return r.contents
}

