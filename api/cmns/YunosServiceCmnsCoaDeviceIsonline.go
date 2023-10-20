package cmns

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// Yunosservicecmnscoadeviceisonline 根据设备id查询设备是否在线
// yunos.service.cmns.coa.device.isonline
//
// 根据设备id查询设备是否在线
func Yunosservicecmnscoadeviceisonline(clt *core.SDKClient, req *cmns.YunosservicecmnscoadeviceisonlineAPIRequest, session string) (*cmns.YunosservicecmnscoadeviceisonlineAPIResponse, error) {
	var resp cmns.YunosservicecmnscoadeviceisonlineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
