package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
access token 获取精灵用户 id APIRequest
alibaba.ailab.user.open.uid.get

access token 获取精灵用户 id
*/
type AlibabaAilabUserOpenUidGetRequest struct {
    model.Params

    // access token
    skillAccessToken   string 

    // skill id
    skillId   int64 

}

func NewAlibabaAilabUserOpenUidGetRequest() *AlibabaAilabUserOpenUidGetRequest{
    return &AlibabaAilabUserOpenUidGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabUserOpenUidGetRequest) GetApiMethodName() string {
    return "alibaba.ailab.user.open.uid.get"
}

func (r AlibabaAilabUserOpenUidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabUserOpenUidGetRequest) SetSkillAccessToken(skillAccessToken string) error {
    r.skillAccessToken = skillAccessToken
    r.Set("skill_access_token", skillAccessToken)
    return nil
}

func (r AlibabaAilabUserOpenUidGetRequest) GetSkillAccessToken() string {
    return r.skillAccessToken
}

func (r *AlibabaAilabUserOpenUidGetRequest) SetSkillId(skillId int64) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

func (r AlibabaAilabUserOpenUidGetRequest) GetSkillId() int64 {
    return r.skillId
}

