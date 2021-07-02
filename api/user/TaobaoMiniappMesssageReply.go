package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoMiniappMesssageReply 轻店铺下行回复接口
// taobao.miniapp.messsage.reply
//
// 外部 isv 调用该进口来进行轻店铺消息的回复
func TaobaoMiniappMesssageReply(clt *core.SDKClient, req *user.TaobaoMiniappMesssageReplyAPIRequest, session string) (*user.TaobaoMiniappMesssageReplyAPIResponse, error) {
	var resp user.TaobaoMiniappMesssageReplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
