package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarFinanceStatusSync 汽车金融状态同步
// tmall.car.finance.status.sync
//
// 汽车金融状态同步
func TmallCarFinanceStatusSync(clt *core.SDKClient, req *tmallcar.TmallCarFinanceStatusSyncAPIRequest, resp *tmallcar.TmallCarFinanceStatusSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
