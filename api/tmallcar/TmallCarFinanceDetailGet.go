package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarFinanceDetailGet 查询汽车金融订单信息
// tmall.car.finance.detail.get
//
// 查询汽车金融订单信息
func TmallCarFinanceDetailGet(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallCarFinanceDetailGetAPIRequest, resp *tmallcar.TmallCarFinanceDetailGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
