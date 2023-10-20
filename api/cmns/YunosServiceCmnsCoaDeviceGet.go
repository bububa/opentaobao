package cmns

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// Yunosservicecmnscoadeviceget 设备详情查询
// yunos.service.cmns.coa.device.get
//
// 第三方应用开发者调用此接口查询设备详情
func Yunosservicecmnscoadeviceget(clt *core.SDKClient, req *cmns.YunosservicecmnscoadevicegetAPIRequest, session string) (*cmns.YunosservicecmnscoadevicegetAPIResponse, error) {
	var resp cmns.YunosservicecmnscoadevicegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
