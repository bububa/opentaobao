package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaooctradesbytagget 标签查询订单
// taobao.oc.trades.bytag.get
//
// 根据标签查询订单编号
func Taobaooctradesbytagget(clt *core.SDKClient, req *jst.TaobaooctradesbytaggetAPIRequest, session string) (*jst.TaobaooctradesbytaggetAPIResponse, error) {
	var resp jst.TaobaooctradesbytaggetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
