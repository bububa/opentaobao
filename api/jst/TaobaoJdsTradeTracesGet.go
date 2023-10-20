package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojdstradetracesget 获取单条订单跟踪详情
// taobao.jds.trade.traces.get
//
// 获取聚石塔数据共享的交易全链路信息
func Taobaojdstradetracesget(clt *core.SDKClient, req *jst.TaobaojdstradetracesgetAPIRequest, session string) (*jst.TaobaojdstradetracesgetAPIResponse, error) {
	var resp jst.TaobaojdstradetracesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
