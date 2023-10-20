package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoyphrefundget 一盘货商家单个查询退款单信息
// taobao.fenxiao.yph.refund.get
//
// 一盘货商家单个查询退款单信息
func Taobaofenxiaoyphrefundget(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoyphrefundgetAPIRequest, session string) (*fenxiao.TaobaofenxiaoyphrefundgetAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoyphrefundgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
