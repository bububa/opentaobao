package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderstorerefused Pos端门店拒单
// taobao.omniorder.store.refused
//
// ISV Pos端门店拒单，通知星盘
func Taobaoomniorderstorerefused(clt *core.SDKClient, req *omniorder.TaobaoomniorderstorerefusedAPIRequest, session string) (*omniorder.TaobaoomniorderstorerefusedAPIResponse, error) {
	var resp omniorder.TaobaoomniorderstorerefusedAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
