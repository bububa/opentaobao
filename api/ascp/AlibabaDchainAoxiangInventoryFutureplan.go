package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxianginventoryfutureplan 负卖计划
// alibaba.dchain.aoxiang.inventory.futureplan
//
// 负卖计划。底层有白名单控制，并非对所有商家开放。如果需要使用，请联系对应的小二增加白名单
func Alibabadchainaoxianginventoryfutureplan(clt *core.SDKClient, req *ascp.AlibabadchainaoxianginventoryfutureplanAPIRequest, session string) (*ascp.AlibabadchainaoxianginventoryfutureplanAPIResponse, error) {
	var resp ascp.AlibabadchainaoxianginventoryfutureplanAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
