package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpshopcategoryfindlist 人群相关类目查询
// taobao.universalbp.shopcategory.findlist
//
// 查询店铺所属的类目信息
func Taobaouniversalbpshopcategoryfindlist(clt *core.SDKClient, req *simba.TaobaouniversalbpshopcategoryfindlistAPIRequest, session string) (*simba.TaobaouniversalbpshopcategoryfindlistAPIResponse, error) {
	var resp simba.TaobaouniversalbpshopcategoryfindlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
