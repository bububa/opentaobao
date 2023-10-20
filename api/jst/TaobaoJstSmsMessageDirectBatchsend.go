package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsMessageDirectBatchsend OAID批量发送，支持明文手机号发送
// taobao.jst.sms.message.direct.batchsend
//
// OAID批量发送，支持明文手机号发送
func TaobaoJstSmsMessageDirectBatchsend(clt *core.SDKClient, req *jst.TaobaoJstSmsMessageDirectBatchsendAPIRequest, resp *jst.TaobaoJstSmsMessageDirectBatchsendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
