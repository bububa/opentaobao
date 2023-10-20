package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderAftersaleOrderGet 闲鱼帮卖售后服务单查询
// alibaba.idle.tender.aftersale.order.get
//
// 闲鱼帮卖售后服务单查询
func AlibabaIdleTenderAftersaleOrderGet(clt *core.SDKClient, req *idle.AlibabaIdleTenderAftersaleOrderGetAPIRequest, resp *idle.AlibabaIdleTenderAftersaleOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
