package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojdstradesstatisticsdiff 订单全链路状态统计差异比较
// taobao.jds.trades.statistics.diff
//
// 订单全链路状态统计差异比较
func Taobaojdstradesstatisticsdiff(clt *core.SDKClient, req *jst.TaobaojdstradesstatisticsdiffAPIRequest, session string) (*jst.TaobaojdstradesstatisticsdiffAPIResponse, error) {
	var resp jst.TaobaojdstradesstatisticsdiffAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
