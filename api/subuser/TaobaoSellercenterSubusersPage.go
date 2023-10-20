package subuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// Taobaosellercentersubuserspage 通过主账号登陆态分页查询子账号列表
// taobao.sellercenter.subusers.page
//
// 通过主账号登陆态分页查询子账号列表
func Taobaosellercentersubuserspage(clt *core.SDKClient, req *subuser.TaobaosellercentersubuserspageAPIRequest, session string) (*subuser.TaobaosellercentersubuserspageAPIResponse, error) {
	var resp subuser.TaobaosellercentersubuserspageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
