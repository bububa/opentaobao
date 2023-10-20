package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Cainiaoglobalimpickupbigbaginfo 大包状态查询
// cainiao.global.im.pickup.bigbag.info
//
// 大包状态查询
func Cainiaoglobalimpickupbigbaginfo(clt *core.SDKClient, req *wlbimports.CainiaoglobalimpickupbigbaginfoAPIRequest, session string) (*wlbimports.CainiaoglobalimpickupbigbaginfoAPIResponse, error) {
	var resp wlbimports.CainiaoglobalimpickupbigbaginfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
