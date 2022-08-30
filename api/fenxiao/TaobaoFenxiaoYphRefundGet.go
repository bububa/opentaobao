package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoYphRefundGet 一盘货商家单个查询退款单信息
// taobao.fenxiao.yph.refund.get
//
// 一盘货商家单个查询退款单信息
func TaobaoFenxiaoYphRefundGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoYphRefundGetAPIRequest, session string) (*fenxiao.TaobaoFenxiaoYphRefundGetAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoYphRefundGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
