package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderdetail 全渠道订单详情
// taobao.omni.order.detail
//
// 全渠道订单详情
func Taobaoomniorderdetail(clt *core.SDKClient, req *omniorder.TaobaoomniorderdetailAPIRequest, session string) (*omniorder.TaobaoomniorderdetailAPIResponse, error) {
	var resp omniorder.TaobaoomniorderdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
