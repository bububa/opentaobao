package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoyphrefundsget 一盘货商家批量查询退款单信息
// taobao.fenxiao.yph.refunds.get
//
// 一盘货商家批量查询退款单信息
func Taobaofenxiaoyphrefundsget(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoyphrefundsgetAPIRequest, session string) (*fenxiao.TaobaofenxiaoyphrefundsgetAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoyphrefundsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
