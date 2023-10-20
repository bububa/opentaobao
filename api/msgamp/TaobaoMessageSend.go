package msgamp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/msgamp"
)

// Taobaomessagesend 消息发送
// taobao.message.send
//
// 消息发送接口
func Taobaomessagesend(clt *core.SDKClient, req *msgamp.TaobaomessagesendAPIRequest, session string) (*msgamp.TaobaomessagesendAPIResponse, error) {
	var resp msgamp.TaobaomessagesendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
