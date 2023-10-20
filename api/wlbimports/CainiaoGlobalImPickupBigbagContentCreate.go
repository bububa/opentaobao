package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Cainiaoglobalimpickupbigbagcontentcreate 大包创建
// cainiao.global.im.pickup.bigbag.content.create
//
// 大包创建
func Cainiaoglobalimpickupbigbagcontentcreate(clt *core.SDKClient, req *wlbimports.CainiaoglobalimpickupbigbagcontentcreateAPIRequest, session string) (*wlbimports.CainiaoglobalimpickupbigbagcontentcreateAPIResponse, error) {
	var resp wlbimports.CainiaoglobalimpickupbigbagcontentcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
