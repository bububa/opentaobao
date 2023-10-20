package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentdevicegetvendor 查询设备Vendor信息
// yunos.tvpubadmin.content.device.getvendor
//
// 查询设备Vendor信息
func Yunostvpubadmincontentdevicegetvendor(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentdevicegetvendorAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentdevicegetvendorAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentdevicegetvendorAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
