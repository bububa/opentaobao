package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderstoreaccpeted Pos端门店接单接口
// taobao.omniorder.store.accpeted
//
// ISV Pos端门店接单，通知星盘
func Taobaoomniorderstoreaccpeted(clt *core.SDKClient, req *omniorder.TaobaoomniorderstoreaccpetedAPIRequest, session string) (*omniorder.TaobaoomniorderstoreaccpetedAPIResponse, error) {
	var resp omniorder.TaobaoomniorderstoreaccpetedAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
