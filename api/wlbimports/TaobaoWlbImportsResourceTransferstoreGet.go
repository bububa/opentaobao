package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Taobaowlbimportsresourcetransferstoreget 根据指定的资源获取所有中转仓列表
// taobao.wlb.imports.resource.transferstore.get
//
// 根据指定的资源获取所有中转仓列表
func Taobaowlbimportsresourcetransferstoreget(clt *core.SDKClient, req *wlbimports.TaobaowlbimportsresourcetransferstoregetAPIRequest, session string) (*wlbimports.TaobaowlbimportsresourcetransferstoregetAPIResponse, error) {
	var resp wlbimports.TaobaowlbimportsresourcetransferstoregetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
