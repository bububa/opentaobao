package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojdsrefundtracesget 获取单条退款跟踪详情
// taobao.jds.refund.traces.get
//
// 获取聚石塔数据共享的交易全链路的退款信息
func Taobaojdsrefundtracesget(clt *core.SDKClient, req *jst.TaobaojdsrefundtracesgetAPIRequest, session string) (*jst.TaobaojdsrefundtracesgetAPIResponse, error) {
	var resp jst.TaobaojdsrefundtracesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
