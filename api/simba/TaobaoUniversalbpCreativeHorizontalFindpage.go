package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpcreativehorizontalfindpage 横向管理创意分页查询
// taobao.universalbp.creative.horizontal.findpage
//
// 横向管理创意分页查询
func Taobaouniversalbpcreativehorizontalfindpage(clt *core.SDKClient, req *simba.TaobaouniversalbpcreativehorizontalfindpageAPIRequest, session string) (*simba.TaobaouniversalbpcreativehorizontalfindpageAPIResponse, error) {
	var resp simba.TaobaouniversalbpcreativehorizontalfindpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
