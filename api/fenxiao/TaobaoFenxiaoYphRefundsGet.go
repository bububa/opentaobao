package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoYphRefundsGet 一盘货商家批量查询退款单信息
// taobao.fenxiao.yph.refunds.get
//
// 一盘货商家批量查询退款单信息
func TaobaoFenxiaoYphRefundsGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoYphRefundsGetAPIRequest, session string) (*fenxiao.TaobaoFenxiaoYphRefundsGetAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoYphRefundsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
