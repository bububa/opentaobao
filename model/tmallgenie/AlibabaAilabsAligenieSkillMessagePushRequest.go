package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵消息推送标准接口 APIRequest
alibaba.ailabs.aligenie.skill.message.push

用于AliGenie技能开发者在技能内主动向音响推送消息的标准服务接口，只有订阅过该消息的用户才能收到消息；
*/
type AlibabaAilabsAligenieSkillMessagePushRequest struct {
    model.Params

    // 要推送的消息内容
    content   string 

    // 智能应用平台创建的技能id
    skillId   int64 

    // 接收方的用户Id，从技能WebHook中取得的userOpenId
    accountType   string 

    // 消息推送的方式，和技能中申请的权限相关，可选值为TO_USER，TO_APP_BOX，BROADCAST
    pushType   string 

    // 是否是测试消息
    test   bool 

    // TO_USER时必填，接收方的用户Id，从技能WebHook中取得的userOpenId
    userId   string 

    // 接收方的用户设备id，从技能WebHook中取得的deviceOpenId，填写设备id，则用户id必填，否则无法推送
    uuid   string 

    // 鉴权用户类型
    authAccountType   string 

}

func NewAlibabaAilabsAligenieSkillMessagePushRequest() *AlibabaAilabsAligenieSkillMessagePushRequest{
    return &AlibabaAilabsAligenieSkillMessagePushRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.skill.message.push"
}

func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsAligenieSkillMessagePushRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetContent() string {
    return r.content
}

func (r *AlibabaAilabsAligenieSkillMessagePushRequest) SetSkillId(skillId int64) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetSkillId() int64 {
    return r.skillId
}

func (r *AlibabaAilabsAligenieSkillMessagePushRequest) SetAccountType(accountType string) error {
    r.accountType = accountType
    r.Set("account_type", accountType)
    return nil
}

func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetAccountType() string {
    return r.accountType
}

func (r *AlibabaAilabsAligenieSkillMessagePushRequest) SetPushType(pushType string) error {
    r.pushType = pushType
    r.Set("push_type", pushType)
    return nil
}

func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetPushType() string {
    return r.pushType
}

func (r *AlibabaAilabsAligenieSkillMessagePushRequest) SetTest(test bool) error {
    r.test = test
    r.Set("test", test)
    return nil
}

func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetTest() bool {
    return r.test
}

func (r *AlibabaAilabsAligenieSkillMessagePushRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetUserId() string {
    return r.userId
}

func (r *AlibabaAilabsAligenieSkillMessagePushRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetUuid() string {
    return r.uuid
}

func (r *AlibabaAilabsAligenieSkillMessagePushRequest) SetAuthAccountType(authAccountType string) error {
    r.authAccountType = authAccountType
    r.Set("auth_account_type", authAccountType)
    return nil
}

func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetAuthAccountType() string {
    return r.authAccountType
}

