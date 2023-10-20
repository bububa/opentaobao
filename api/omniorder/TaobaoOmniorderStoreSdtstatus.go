package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderstoresdtstatus 菜鸟裹裹运单状态查询
// taobao.omniorder.store.sdtstatus
//
// 提供给商家查询运力单的状态。
func Taobaoomniorderstoresdtstatus(clt *core.SDKClient, req *omniorder.TaobaoomniorderstoresdtstatusAPIRequest, session string) (*omniorder.TaobaoomniorderstoresdtstatusAPIResponse, error) {
	var resp omniorder.TaobaoomniorderstoresdtstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
