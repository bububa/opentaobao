package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// Taobaominiappdistributionorderget 小程序投放-查询小程序投放计划信息
// taobao.miniapp.distribution.order.get
//
// 服务商可通过该API，读取自己开发的小程序对应的投放计划的相关信息
func Taobaominiappdistributionorderget(clt *core.SDKClient, req *miniappopen.TaobaominiappdistributionordergetAPIRequest, session string) (*miniappopen.TaobaominiappdistributionordergetAPIResponse, error) {
	var resp miniappopen.TaobaominiappdistributionordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
