package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmpointruleget 查询积分规则
// alibaba.alsc.crm.point.rule.get
//
// 新增积分规则查询接口,传入includeLogicalDelete和maxUpdateTime时走同步下行逻辑不然则走普通积分规则查询接口
func Alibabaalsccrmpointruleget(clt *core.SDKClient, req *alsc.AlibabaalsccrmpointrulegetAPIRequest, session string) (*alsc.AlibabaalsccrmpointrulegetAPIResponse, error) {
	var resp alsc.AlibabaalsccrmpointrulegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
