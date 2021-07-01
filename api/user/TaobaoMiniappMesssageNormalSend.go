package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

/* TaobaoMiniappMesssageNormalSend
轻店铺下行普通消息给用户
taobao.miniapp.messsage.normal.send

小程序下行单个普通消息 */
func TaobaoMiniappMesssageNormalSend(clt *core.SDKClient, req *user.TaobaoMiniappMesssageNormalSendAPIRequest, session string) (*user.TaobaoMiniappMesssageNormalSendAPIResponse, error) {
	var resp user.TaobaoMiniappMesssageNormalSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
