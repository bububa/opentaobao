package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieSkillMessagePushAPIRequest 天猫精灵消息推送标准接口 API请求
// alibaba.ailabs.aligenie.skill.message.push
//
// 用于AliGenie技能开发者在技能内主动向音响推送消息的标准服务接口，只有订阅过该消息的用户才能收到消息；
type AlibabaAilabsAligenieSkillMessagePushAPIRequest struct {
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

// NewAlibabaAilabsAligenieSkillMessagePushRequest 初始化AlibabaAilabsAligenieSkillMessagePushAPIRequest对象
func NewAlibabaAilabsAligenieSkillMessagePushRequest() *AlibabaAilabsAligenieSkillMessagePushAPIRequest {
	return &AlibabaAilabsAligenieSkillMessagePushAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsAligenieSkillMessagePushAPIRequest) Reset() {
	r._authAccountType = ""
	r._content = ""
	r._accountType = ""
	r._pushType = ""
	r._userId = ""
	r._uuid = ""
	r._skillId = 0
	r._test = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieSkillMessagePushAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.skill.message.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieSkillMessagePushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsAligenieSkillMessagePushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthAccountType is AuthAccountType Setter
// 鉴权用户类型
func (r *AlibabaAilabsAligenieSkillMessagePushAPIRequest) SetAuthAccountType(_authAccountType string) error {
	r._authAccountType = _authAccountType
	r.Set("auth_account_type", _authAccountType)
	return nil
}

// GetAuthAccountType AuthAccountType Getter
func (r AlibabaAilabsAligenieSkillMessagePushAPIRequest) GetAuthAccountType() string {
	return r._authAccountType
}

// SetContent is Content Setter
// 要推送的消息内容
func (r *AlibabaAilabsAligenieSkillMessagePushAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabaAilabsAligenieSkillMessagePushAPIRequest) GetContent() string {
	return r._content
}

// SetAccountType is AccountType Setter
// 接收方的用户Id，从技能WebHook中取得的userOpenId
func (r *AlibabaAilabsAligenieSkillMessagePushAPIRequest) SetAccountType(_accountType string) error {
	r._accountType = _accountType
	r.Set("account_type", _accountType)
	return nil
}

// GetAccountType AccountType Getter
func (r AlibabaAilabsAligenieSkillMessagePushAPIRequest) GetAccountType() string {
	return r._accountType
}

// SetPushType is PushType Setter
// 消息推送的方式，和技能中申请的权限相关，可选值为TO_USER，TO_APP_BOX，BROADCAST
func (r *AlibabaAilabsAligenieSkillMessagePushAPIRequest) SetPushType(_pushType string) error {
	r._pushType = _pushType
	r.Set("push_type", _pushType)
	return nil
}

// GetPushType PushType Getter
func (r AlibabaAilabsAligenieSkillMessagePushAPIRequest) GetPushType() string {
	return r._pushType
}

// SetUserId is UserId Setter
// TO_USER时必填，接收方的用户Id，从技能WebHook中取得的userOpenId
func (r *AlibabaAilabsAligenieSkillMessagePushAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaAilabsAligenieSkillMessagePushAPIRequest) GetUserId() string {
	return r._userId
}

// SetUuid is Uuid Setter
// 接收方的用户设备id，从技能WebHook中取得的deviceOpenId，填写设备id，则用户id必填，否则无法推送
func (r *AlibabaAilabsAligenieSkillMessagePushAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaAilabsAligenieSkillMessagePushAPIRequest) GetUuid() string {
	return r._uuid
}

// SetSkillId is SkillId Setter
// 智能应用平台创建的技能id
func (r *AlibabaAilabsAligenieSkillMessagePushAPIRequest) SetSkillId(_skillId int64) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r AlibabaAilabsAligenieSkillMessagePushAPIRequest) GetSkillId() int64 {
	return r._skillId
}

// SetTest is Test Setter
// 是否是测试消息
func (r *AlibabaAilabsAligenieSkillMessagePushAPIRequest) SetTest(_test bool) error {
	r._test = _test
	r.Set("test", _test)
	return nil
}

// GetTest Test Getter
func (r AlibabaAilabsAligenieSkillMessagePushAPIRequest) GetTest() bool {
	return r._test
}

var poolAlibabaAilabsAligenieSkillMessagePushAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsAligenieSkillMessagePushRequest()
	},
}

// GetAlibabaAilabsAligenieSkillMessagePushRequest 从 sync.Pool 获取 AlibabaAilabsAligenieSkillMessagePushAPIRequest
func GetAlibabaAilabsAligenieSkillMessagePushAPIRequest() *AlibabaAilabsAligenieSkillMessagePushAPIRequest {
	return poolAlibabaAilabsAligenieSkillMessagePushAPIRequest.Get().(*AlibabaAilabsAligenieSkillMessagePushAPIRequest)
}

// ReleaseAlibabaAilabsAligenieSkillMessagePushAPIRequest 将 AlibabaAilabsAligenieSkillMessagePushAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsAligenieSkillMessagePushAPIRequest(v *AlibabaAilabsAligenieSkillMessagePushAPIRequest) {
	v.Reset()
	poolAlibabaAilabsAligenieSkillMessagePushAPIRequest.Put(v)
}
