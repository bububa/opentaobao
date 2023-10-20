package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoyphordersget 批量查询一盘货采购单信息
// taobao.fenxiao.yph.orders.get
//
// 一盘货商家批量查询采购单信息
func Taobaofenxiaoyphordersget(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoyphordersgetAPIRequest, session string) (*fenxiao.TaobaofenxiaoyphordersgetAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoyphordersgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
