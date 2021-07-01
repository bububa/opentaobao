package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieSkillMessagePushAPIRequest
天猫精灵消息推送标准接口 API请求
alibaba.ailabs.aligenie.skill.message.push

用于AliGenie技能开发者在技能内主动向音响推送消息的标准服务接口，只有订阅过该消息的用户才能收到消息； */
type AlibabaAilabsAligenieSkillMessagePushAPIRequest struct {
	model.Params
	// 要推送的消息内容
	_content string
	// 智能应用平台创建的技能id
	_skillId int64
	// 接收方的用户Id，从技能WebHook中取得的userOpenId
	_accountType string
	// 消息推送的方式，和技能中申请的权限相关，可选值为TO_USER，TO_APP_BOX，BROADCAST
	_pushType string
	// 是否是测试消息
	_test bool
	// TO_USER时必填，接收方的用户Id，从技能WebHook中取得的userOpenId
	_userId string
	// 接收方的用户设备id，从技能WebHook中取得的deviceOpenId，填写设备id，则用户id必填，否则无法推送
	_uuid string
	// 鉴权用户类型
	_authAccountType string
}

// New
