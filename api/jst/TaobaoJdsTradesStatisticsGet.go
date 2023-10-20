package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojdstradesstatisticsget 获取订单数量统计结果
// taobao.jds.trades.statistics.get
//
// 获取订单数量统计结果
func Taobaojdstradesstatisticsget(clt *core.SDKClient, req *jst.TaobaojdstradesstatisticsgetAPIRequest, session string) (*jst.TaobaojdstradesstatisticsgetAPIResponse, error) {
	var resp jst.TaobaojdstradesstatisticsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
