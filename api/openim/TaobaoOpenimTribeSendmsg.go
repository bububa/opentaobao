package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimtribesendmsg 发送群消息
// taobao.openim.tribe.sendmsg
//
// 发送群消息，目前支持发送4种类型的群消息，普通文本，图片，语音，自定义消息
func Taobaoopenimtribesendmsg(clt *core.SDKClient, req *openim.TaobaoopenimtribesendmsgAPIRequest, session string) (*openim.TaobaoopenimtribesendmsgAPIResponse, error) {
	var resp openim.TaobaoopenimtribesendmsgAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
