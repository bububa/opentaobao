package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojstsmsmessagesend 聚石塔数据paas短信发送接口
// taobao.jst.sms.message.send
//
// 聚石塔短信PAAS场景中，ISV通过该API帮商家发送短信给用户。
func Taobaojstsmsmessagesend(clt *core.SDKClient, req *jst.TaobaojstsmsmessagesendAPIRequest, session string) (*jst.TaobaojstsmsmessagesendAPIResponse, error) {
	var resp jst.TaobaojstsmsmessagesendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
