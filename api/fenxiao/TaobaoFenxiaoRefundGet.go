package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaorefundget 查询采购单退款信息
// taobao.fenxiao.refund.get
//
// 分销商或供应商可以查询某子单的退款信息，以及下游订单的退款信息
func Taobaofenxiaorefundget(clt *core.SDKClient, req *fenxiao.TaobaofenxiaorefundgetAPIRequest, session string) (*fenxiao.TaobaofenxiaorefundgetAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaorefundgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
