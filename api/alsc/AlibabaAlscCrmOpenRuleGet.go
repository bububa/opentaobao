package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmopenruleget 查询规则
// alibaba.alsc.crm.open.rule.get
//
// 查询会员规则
func Alibabaalsccrmopenruleget(clt *core.SDKClient, req *alsc.AlibabaalsccrmopenrulegetAPIRequest, session string) (*alsc.AlibabaalsccrmopenrulegetAPIResponse, error) {
	var resp alsc.AlibabaalsccrmopenrulegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
