package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAiUserQuickRegister 精灵用户注册申请
// alibaba.ai.user.quick.register
//
// 人工智能实验室精灵用户注册申请接口，开放给Iot厂商做厂商会员数据上报
func AlibabaAiUserQuickRegister(clt *core.SDKClient, req *tmallgenie.AlibabaAiUserQuickRegisterAPIRequest, resp *tmallgenie.AlibabaAiUserQuickRegisterAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
