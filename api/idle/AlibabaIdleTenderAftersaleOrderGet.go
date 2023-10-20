package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidletenderaftersaleorderget 闲鱼帮卖售后服务单查询
// alibaba.idle.tender.aftersale.order.get
//
// 闲鱼帮卖售后服务单查询
func Alibabaidletenderaftersaleorderget(clt *core.SDKClient, req *idle.AlibabaidletenderaftersaleordergetAPIRequest, session string) (*idle.AlibabaidletenderaftersaleordergetAPIResponse, error) {
	var resp idle.AlibabaidletenderaftersaleordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
