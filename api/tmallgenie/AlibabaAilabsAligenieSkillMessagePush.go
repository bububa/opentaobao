package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsAligenieSkillMessagePush 天猫精灵消息推送标准接口
// alibaba.ailabs.aligenie.skill.message.push
//
// 用于AliGenie技能开发者在技能内主动向音响推送消息的标准服务接口，只有订阅过该消息的用户才能收到消息；
func AlibabaAilabsAligenieSkillMessagePush(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsAligenieSkillMessagePushAPIRequest, session string) (*tmallgenie.AlibabaAilabsAligenieSkillMessagePushAPIResponse, error) {
	var resp tmallgenie.AlibabaAilabsAligenieSkillMessagePushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
