package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsaligenieskillmessagepushAPIRequest 天猫精灵消息推送标准接口 API请求
// alibaba.ailabs.aligenie.skill.message.push
//
// 用于AliGenie技能开发者在技能内主动向音响推送消息的标准服务接口，只有订阅过该消息的用户才能收到消息；
type AlibabaailabsaligenieskillmessagepushAPIRequest struct {
	model.Params
	// 鉴权用户类型
	_authAccountType string
	// 要推送的消息内容
	_content string
	// 接收方的用户Id，从技能WebHook中取得的userOpenId
	_accountType string
	// 消息推送的方式，和技能中申请的权限相关，可选值为TO_USER，TO_APP_BOX，BROADCAST
	_pushType string
	// TO_USER时必填，接收方的用户Id，从技能WebHook中取得的userOpenId
	_userId string
	// 接收方的用户设备id，从技能WebHook中取得的deviceOpenId，填写设备id，则用户id必填，否则无法推送
	_uuid string
	// 智能应用平台创建的技能id
	_skillId int64
	// 是否是测试消息
	_test bool
}

// NewAlibabaailabsaligenieskillmessagepushRequest 初始化AlibabaailabsaligenieskillmessagepushAPIRequest对象
func NewAlibabaailabsaligenieskillmessagepushRequest() *AlibabaailabsaligenieskillmessagepushAPIRequest {
	return &AlibabaailabsaligenieskillmessagepushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsaligenieskillmessagepushAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.skill.message.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsaligenieskillmessagepushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsaligenieskillmessagepushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthAccountType is AuthAccountType Setter
// 鉴权用户类型
func (r *AlibabaailabsaligenieskillmessagepushAPIRequest) SetAuthAccountType(_authAccountType string) error {
	r._authAccountType = _authAccountType
	r.Set("auth_account_type", _authAccountType)
	return nil
}

// GetAuthAccountType AuthAccountType Getter
func (r AlibabaailabsaligenieskillmessagepushAPIRequest) GetAuthAccountType() string {
	return r._authAccountType
}

// SetContent is Content Setter
// 要推送的消息内容
func (r *AlibabaailabsaligenieskillmessagepushAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabaailabsaligenieskillmessagepushAPIRequest) GetContent() string {
	return r._content
}

// SetAccountType is AccountType Setter
// 接收方的用户Id，从技能WebHook中取得的userOpenId
func (r *AlibabaailabsaligenieskillmessagepushAPIRequest) SetAccountType(_accountType string) error {
	r._accountType = _accountType
	r.Set("account_type", _accountType)
	return nil
}

// GetAccountType AccountType Getter
func (r AlibabaailabsaligenieskillmessagepushAPIRequest) GetAccountType() string {
	return r._accountType
}

// SetPushType is PushType Setter
// 消息推送的方式，和技能中申请的权限相关，可选值为TO_USER，TO_APP_BOX，BROADCAST
func (r *AlibabaailabsaligenieskillmessagepushAPIRequest) SetPushType(_pushType string) error {
	r._pushType = _pushType
	r.Set("push_type", _pushType)
	return nil
}

// GetPushType PushType Getter
func (r AlibabaailabsaligenieskillmessagepushAPIRequest) GetPushType() string {
	return r._pushType
}

// SetUserId is UserId Setter
// TO_USER时必填，接收方的用户Id，从技能WebHook中取得的userOpenId
func (r *AlibabaailabsaligenieskillmessagepushAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaailabsaligenieskillmessagepushAPIRequest) GetUserId() string {
	return r._userId
}

// SetUuid is Uuid Setter
// 接收方的用户设备id，从技能WebHook中取得的deviceOpenId，填写设备id，则用户id必填，否则无法推送
func (r *AlibabaailabsaligenieskillmessagepushAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaailabsaligenieskillmessagepushAPIRequest) GetUuid() string {
	return r._uuid
}

// SetSkillId is SkillId Setter
// 智能应用平台创建的技能id
func (r *AlibabaailabsaligenieskillmessagepushAPIRequest) SetSkillId(_skillId int64) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r AlibabaailabsaligenieskillmessagepushAPIRequest) GetSkillId() int64 {
	return r._skillId
}

// SetTest is Test Setter
// 是否是测试消息
func (r *AlibabaailabsaligenieskillmessagepushAPIRequest) SetTest(_test bool) error {
	r._test = _test
	r.Set("test", _test)
	return nil
}

// GetTest Test Getter
func (r AlibabaailabsaligenieskillmessagepushAPIRequest) GetTest() bool {
	return r._test
}
