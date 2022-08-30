package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarFinanceStatusSync 汽车金融状态同步
// tmall.car.finance.status.sync
//
// 汽车金融状态同步
func TmallCarFinanceStatusSync(clt *core.SDKClient, req *tmallcar.TmallCarFinanceStatusSyncAPIRequest, session string) (*tmallcar.TmallCarFinanceStatusSyncAPIResponse, error) {
	var resp tmallcar.TmallCarFinanceStatusSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
