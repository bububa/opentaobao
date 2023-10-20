package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxytriggerevent 抽奖活动自定义事件触发
// alitrip.merchant.galaxy.trigger.event
//
// 自定义事件触发
func Alitripmerchantgalaxytriggerevent(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxytriggereventAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxytriggereventAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxytriggereventAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
