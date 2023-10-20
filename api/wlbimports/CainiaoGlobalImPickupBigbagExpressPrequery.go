package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Cainiaoglobalimpickupbigbagexpressprequery 首公里揽收-快递预查询服务
// cainiao.global.im.pickup.bigbag.express.prequery
//
// 快递预查询服务
func Cainiaoglobalimpickupbigbagexpressprequery(clt *core.SDKClient, req *wlbimports.CainiaoglobalimpickupbigbagexpressprequeryAPIRequest, session string) (*wlbimports.CainiaoglobalimpickupbigbagexpressprequeryAPIResponse, error) {
	var resp wlbimports.CainiaoglobalimpickupbigbagexpressprequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
