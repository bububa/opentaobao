package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmExchangeCrowdinstanceAddPrivy 向活动人群实例中增加买家（隐私号版）
// taobao.crm.exchange.crowdinstance.add.privy
//
// 向活动人群实例中增加买家
func TaobaoCrmExchangeCrowdinstanceAddPrivy(clt *core.SDKClient, req *crm.TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest, session string) (*crm.TaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse, error) {
	var resp crm.TaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
