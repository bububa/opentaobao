package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Cainiaoglobalimpickupstoresget 首公里揽收-集货仓列表查询
// cainiao.global.im.pickup.stores.get
//
// 首公里揽收-集货仓列表查询
func Cainiaoglobalimpickupstoresget(clt *core.SDKClient, req *wlbimports.CainiaoglobalimpickupstoresgetAPIRequest, session string) (*wlbimports.CainiaoglobalimpickupstoresgetAPIResponse, error) {
	var resp wlbimports.CainiaoglobalimpickupstoresgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
