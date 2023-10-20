package msgamp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/msgamp"
)

// Taobaobcchatmessagesend 小程序资源授权-BC客服消息
// taobao.bc.chat.message.send
//
// 小程序资源授权-消息订阅
func Taobaobcchatmessagesend(clt *core.SDKClient, req *msgamp.TaobaobcchatmessagesendAPIRequest, session string) (*msgamp.TaobaobcchatmessagesendAPIResponse, error) {
	var resp msgamp.TaobaobcchatmessagesendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
