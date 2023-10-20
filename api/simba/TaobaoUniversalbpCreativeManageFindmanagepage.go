package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpcreativemanagefindmanagepage 创意库查询创意列表
// taobao.universalbp.creative.manage.findmanagepage
//
// 创意库查询创意列表
func Taobaouniversalbpcreativemanagefindmanagepage(clt *core.SDKClient, req *simba.TaobaouniversalbpcreativemanagefindmanagepageAPIRequest, session string) (*simba.TaobaouniversalbpcreativemanagefindmanagepageAPIResponse, error) {
	var resp simba.TaobaouniversalbpcreativemanagefindmanagepageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
