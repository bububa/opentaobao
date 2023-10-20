package gameact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/gameact"
)

// Taobaodeactivitydeliveryaddrconfirm 用户收件地址确认
// taobao.de.activity.delivery.addr.confirm
//
// 用户收件地址确认
func Taobaodeactivitydeliveryaddrconfirm(clt *core.SDKClient, req *gameact.TaobaodeactivitydeliveryaddrconfirmAPIRequest, session string) (*gameact.TaobaodeactivitydeliveryaddrconfirmAPIResponse, error) {
	var resp gameact.TaobaodeactivitydeliveryaddrconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
