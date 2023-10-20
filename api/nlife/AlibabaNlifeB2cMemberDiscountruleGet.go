package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// Alibabanlifeb2cmemberdiscountruleget 会员抵扣规则
// alibaba.nlife.b2c.member.discountrule.get
//
// 获取企业会员抵扣规则
func Alibabanlifeb2cmemberdiscountruleget(clt *core.SDKClient, req *nlife.Alibabanlifeb2cmemberdiscountrulegetAPIRequest, session string) (*nlife.Alibabanlifeb2cmemberdiscountrulegetAPIResponse, error) {
	var resp nlife.Alibabanlifeb2cmemberdiscountrulegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
