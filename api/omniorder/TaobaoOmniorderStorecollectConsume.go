package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderstorecollectconsume 全渠道门店自提核销订单
// taobao.omniorder.storecollect.consume
//
// 全渠道门店自提核销订单
func Taobaoomniorderstorecollectconsume(clt *core.SDKClient, req *omniorder.TaobaoomniorderstorecollectconsumeAPIRequest, session string) (*omniorder.TaobaoomniorderstorecollectconsumeAPIResponse, error) {
	var resp omniorder.TaobaoomniorderstorecollectconsumeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
