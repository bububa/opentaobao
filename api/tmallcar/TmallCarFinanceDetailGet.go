package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarFinanceDetailGet 查询汽车金融订单信息
// tmall.car.finance.detail.get
//
// 查询汽车金融订单信息
func TmallCarFinanceDetailGet(clt *core.SDKClient, req *tmallcar.TmallCarFinanceDetailGetAPIRequest, session string) (*tmallcar.TmallCarFinanceDetailGetAPIResponse, error) {
	var resp tmallcar.TmallCarFinanceDetailGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
