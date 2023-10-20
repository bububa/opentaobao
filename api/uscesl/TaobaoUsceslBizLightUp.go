package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// Taobaousceslbizlightup 价签LED等点亮
// taobao.uscesl.biz.light.up
//
// 价签LED等点亮
func Taobaousceslbizlightup(clt *core.SDKClient, req *uscesl.TaobaousceslbizlightupAPIRequest, session string) (*uscesl.TaobaousceslbizlightupAPIResponse, error) {
	var resp uscesl.TaobaousceslbizlightupAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
