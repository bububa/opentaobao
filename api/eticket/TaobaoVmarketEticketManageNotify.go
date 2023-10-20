package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaovmarketeticketmanagenotify 主动发起通知接口
// taobao.vmarket.eticket.manage.notify
//
// 外部合作商家主动发起通知接口
func Taobaovmarketeticketmanagenotify(clt *core.SDKClient, req *eticket.TaobaovmarketeticketmanagenotifyAPIRequest, session string) (*eticket.TaobaovmarketeticketmanagenotifyAPIResponse, error) {
	var resp eticket.TaobaovmarketeticketmanagenotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
