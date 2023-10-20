package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaaiuserquickbind 精灵用户绑定第三方账号信息
// alibaba.ai.user.quick.bind
//
// 人工智能实验室精灵用户绑定第三方账号信息接口，开放给Iot厂商做为厂商上送第三方账号信息的接口
func Alibabaaiuserquickbind(clt *core.SDKClient, req *tmallgenie.AlibabaaiuserquickbindAPIRequest, session string) (*tmallgenie.AlibabaaiuserquickbindAPIResponse, error) {
	var resp tmallgenie.AlibabaaiuserquickbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
