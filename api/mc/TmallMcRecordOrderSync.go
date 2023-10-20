package mc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mc"
)

// Tmallmcrecordordersync 订单信息同步
// tmall.mc.record.order.sync
//
// 订单信息同步(零售云接口)
func Tmallmcrecordordersync(clt *core.SDKClient, req *mc.TmallmcrecordordersyncAPIRequest, session string) (*mc.TmallmcrecordordersyncAPIResponse, error) {
	var resp mc.TmallmcrecordordersyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
