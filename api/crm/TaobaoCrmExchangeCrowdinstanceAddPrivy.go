package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// Taobaocrmexchangecrowdinstanceaddprivy 向活动人群实例中增加买家（隐私号版）
// taobao.crm.exchange.crowdinstance.add.privy
//
// 向活动人群实例中增加买家
func Taobaocrmexchangecrowdinstanceaddprivy(clt *core.SDKClient, req *crm.TaobaocrmexchangecrowdinstanceaddprivyAPIRequest, session string) (*crm.TaobaocrmexchangecrowdinstanceaddprivyAPIResponse, error) {
	var resp crm.TaobaocrmexchangecrowdinstanceaddprivyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
