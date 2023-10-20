package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// Taobaodaogoubaoorderstatisticstotal 销售订单总额统计
// taobao.daogoubao.order.statistics.total
//
// 对接千牛端数字中心
func Taobaodaogoubaoorderstatisticstotal(clt *core.SDKClient, req *qianniu.TaobaodaogoubaoorderstatisticstotalAPIRequest, session string) (*qianniu.TaobaodaogoubaoorderstatisticstotalAPIResponse, error) {
	var resp qianniu.TaobaodaogoubaoorderstatisticstotalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
