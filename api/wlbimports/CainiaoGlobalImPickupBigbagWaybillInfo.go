package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Cainiaoglobalimpickupbigbagwaybillinfo 大包面单查询
// cainiao.global.im.pickup.bigbag.waybill.info
//
// 大包面单查询
func Cainiaoglobalimpickupbigbagwaybillinfo(clt *core.SDKClient, req *wlbimports.CainiaoglobalimpickupbigbagwaybillinfoAPIRequest, session string) (*wlbimports.CainiaoglobalimpickupbigbagwaybillinfoAPIResponse, error) {
	var resp wlbimports.CainiaoglobalimpickupbigbagwaybillinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
