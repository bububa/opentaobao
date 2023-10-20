package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Taobaowlbimportsresourceget 获取所有服务列表
// taobao.wlb.imports.resource.get
//
// 一般进口TOP接口，获取所有服务列表。
func Taobaowlbimportsresourceget(clt *core.SDKClient, req *wlbimports.TaobaowlbimportsresourcegetAPIRequest, session string) (*wlbimports.TaobaowlbimportsresourcegetAPIResponse, error) {
	var resp wlbimports.TaobaowlbimportsresourcegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
