package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaaiuserquicktokenbind 人工智能实验室精灵用户绑定第三方Token接口
// alibaba.ai.user.quick.token.bind
//
// 人工智能实验室精灵用户绑定第三方Token接口
func Alibabaaiuserquicktokenbind(clt *core.SDKClient, req *tmallgenie.AlibabaaiuserquicktokenbindAPIRequest, session string) (*tmallgenie.AlibabaaiuserquicktokenbindAPIResponse, error) {
	var resp tmallgenie.AlibabaaiuserquicktokenbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
