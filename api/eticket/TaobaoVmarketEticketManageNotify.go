package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketManageNotify 主动发起通知接口
// taobao.vmarket.eticket.manage.notify
//
// 外部合作商家主动发起通知接口
func TaobaoVmarketEticketManageNotify(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketManageNotifyAPIRequest, session string) (*eticket.TaobaoVmarketEticketManageNotifyAPIResponse, error) {
	var resp eticket.TaobaoVmarketEticketManageNotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
