package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

/* AlibabaAiUserQuickBind
精灵用户绑定第三方账号信息
alibaba.ai.user.quick.bind

人工智能实验室精灵用户绑定第三方账号信息接口，开放给Iot厂商做为厂商上送第三方账号信息的接口 */
func AlibabaAiUserQuickBind(clt *core.SDKClient, req *tmallgenie.AlibabaAiUserQuickBindAPIRequest, session string) (*tmallgenie.AlibabaAiUserQuickBindAPIResponse, error) {
	var resp tmallgenie.AlibabaAiUserQuickBindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
