package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// Taobaousceslbizitemlightup 商品条码亮灯API
// taobao.uscesl.biz.item.light.up
//
// 亮灯API
func Taobaousceslbizitemlightup(clt *core.SDKClient, req *uscesl.TaobaousceslbizitemlightupAPIRequest, session string) (*uscesl.TaobaousceslbizitemlightupAPIResponse, error) {
	var resp uscesl.TaobaousceslbizitemlightupAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
