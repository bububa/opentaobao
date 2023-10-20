package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderAftersaleOrderGet 闲鱼帮卖售后服务单查询
// alibaba.idle.tender.aftersale.order.get
//
// 闲鱼帮卖售后服务单查询
func AlibabaIdleTenderAftersaleOrderGet(clt *core.SDKClient, req *idle.AlibabaIdleTenderAftersaleOrderGetAPIRequest, session string) (*idle.AlibabaIdleTenderAftersaleOrderGetAPIResponse, error) {
	var resp idle.AlibabaIdleTenderAftersaleOrderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
