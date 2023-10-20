package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkposafterbuybenefitsend 生态pos购后发放权益
// alibaba.wdk.pos.afterbuy.benefit.send
//
// 生态pos购后发放权益接口开放
func Alibabawdkposafterbuybenefitsend(clt *core.SDKClient, req *wdk.AlibabawdkposafterbuybenefitsendAPIRequest, session string) (*wdk.AlibabawdkposafterbuybenefitsendAPIResponse, error) {
	var resp wdk.AlibabawdkposafterbuybenefitsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
