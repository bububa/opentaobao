package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallcarfinancedetailget 查询汽车金融订单信息
// tmall.car.finance.detail.get
//
// 查询汽车金融订单信息
func Tmallcarfinancedetailget(clt *core.SDKClient, req *tmallcar.TmallcarfinancedetailgetAPIRequest, session string) (*tmallcar.TmallcarfinancedetailgetAPIResponse, error) {
	var resp tmallcar.TmallcarfinancedetailgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
