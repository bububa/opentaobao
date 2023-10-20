package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoYphOrdersGet 批量查询一盘货采购单信息
// taobao.fenxiao.yph.orders.get
//
// 一盘货商家批量查询采购单信息
func TaobaoFenxiaoYphOrdersGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoYphOrdersGetAPIRequest, session string) (*fenxiao.TaobaoFenxiaoYphOrdersGetAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoYphOrdersGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
