package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmExchangeCrowdinstanceAdd 向活动人群实例中增加买家
// taobao.crm.exchange.crowdinstance.add
//
// 向活动人群实例中增加买家
func TaobaoCrmExchangeCrowdinstanceAdd(clt *core.SDKClient, req *crm.TaobaoCrmExchangeCrowdinstanceAddAPIRequest, session string) (*crm.TaobaoCrmExchangeCrowdinstanceAddAPIResponse, error) {
	var resp crm.TaobaoCrmExchangeCrowdinstanceAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
