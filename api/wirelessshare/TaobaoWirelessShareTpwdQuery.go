package wirelessshare

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wirelessshare"
)

// Taobaowirelesssharetpwdquery 查询解析淘口令
// taobao.wireless.share.tpwd.query
//
// 查询解析淘口令
func Taobaowirelesssharetpwdquery(clt *core.SDKClient, req *wirelessshare.TaobaowirelesssharetpwdqueryAPIRequest, session string) (*wirelessshare.TaobaowirelesssharetpwdqueryAPIResponse, error) {
	var resp wirelessshare.TaobaowirelesssharetpwdqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
