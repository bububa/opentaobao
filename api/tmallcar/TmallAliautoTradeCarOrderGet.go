package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautotradecarorderget 整车订单详情查询
// tmall.aliauto.trade.car.order.get
//
// 整车订单详情查询接口
func Tmallaliautotradecarorderget(clt *core.SDKClient, req *tmallcar.TmallaliautotradecarordergetAPIRequest, session string) (*tmallcar.TmallaliautotradecarordergetAPIResponse, error) {
	var resp tmallcar.TmallaliautotradecarordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
