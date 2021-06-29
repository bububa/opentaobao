package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵消息推送标准接口 API请求
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

// 初始化AlibabaAilabsAligenieSkillMessagePushRequest对象
func NewAlibabaAilabsAligenieSkillMessagePushRequest() *AlibabaAilabsAligenieSkillMessagePushRequest{
    return &AlibabaAilabsAligenieSkillMessagePushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.skill.message.push"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Content Setter
// 要推送的消息内容
func (r *AlibabaAilabsAligenieSkillMessagePushRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetContent() string {
    return r.content
}
// SkillId Setter
// 智能应用平台创建的技能id
func (r *AlibabaAilabsAligenieSkillMessagePushRequest) SetSkillId(skillId int64) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

// SkillId Getter
func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetSkillId() int64 {
    return r.skillId
}
// AccountType Setter
// 接收方的用户Id，从技能WebHook中取得的userOpenId
func (r *AlibabaAilabsAligenieSkillMessagePushRequest) SetAccountType(accountType string) error {
    r.accountType = accountType
    r.Set("account_type", accountType)
    return nil
}

// AccountType Getter
func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetAccountType() string {
    return r.accountType
}
// PushType Setter
// 消息推送的方式，和技能中申请的权限相关，可选值为TO_USER，TO_APP_BOX，BROADCAST
func (r *AlibabaAilabsAligenieSkillMessagePushRequest) SetPushType(pushType string) error {
    r.pushType = pushType
    r.Set("push_type", pushType)
    return nil
}

// PushType Getter
func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetPushType() string {
    return r.pushType
}
// Test Setter
// 是否是测试消息
func (r *AlibabaAilabsAligenieSkillMessagePushRequest) SetTest(test bool) error {
    r.test = test
    r.Set("test", test)
    return nil
}

// Test Getter
func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetTest() bool {
    return r.test
}
// UserId Setter
// TO_USER时必填，接收方的用户Id，从技能WebHook中取得的userOpenId
func (r *AlibabaAilabsAligenieSkillMessagePushRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetUserId() string {
    return r.userId
}
// Uuid Setter
// 接收方的用户设备id，从技能WebHook中取得的deviceOpenId，填写设备id，则用户id必填，否则无法推送
func (r *AlibabaAilabsAligenieSkillMessagePushRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetUuid() string {
    return r.uuid
}
// AuthAccountType Setter
// 鉴权用户类型
func (r *AlibabaAilabsAligenieSkillMessagePushRequest) SetAuthAccountType(authAccountType string) error {
    r.authAccountType = authAccountType
    r.Set("auth_account_type", authAccountType)
    return nil
}

// AuthAccountType Getter
func (r AlibabaAilabsAligenieSkillMessagePushRequest) GetAuthAccountType() string {
    return r.authAccountType
}
