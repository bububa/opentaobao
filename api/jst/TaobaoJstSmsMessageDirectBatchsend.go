package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojstsmsmessagedirectbatchsend OAID批量发送，支持明文手机号发送
// taobao.jst.sms.message.direct.batchsend
//
// OAID批量发送，支持明文手机号发送
func Taobaojstsmsmessagedirectbatchsend(clt *core.SDKClient, req *jst.TaobaojstsmsmessagedirectbatchsendAPIRequest, session string) (*jst.TaobaojstsmsmessagedirectbatchsendAPIResponse, error) {
	var resp jst.TaobaojstsmsmessagedirectbatchsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
