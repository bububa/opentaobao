package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaomessageaccountmesssagenormalsend 下行普通消息
// taobao.messageaccount.messsage.normal.send
//
// 消息号下行单个普通消息
func Taobaomessageaccountmesssagenormalsend(clt *core.SDKClient, req *user.TaobaomessageaccountmesssagenormalsendAPIRequest, session string) (*user.TaobaomessageaccountmesssagenormalsendAPIResponse, error) {
	var resp user.TaobaomessageaccountmesssagenormalsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
