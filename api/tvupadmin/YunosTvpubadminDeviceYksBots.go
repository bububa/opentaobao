package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindeviceyksbots 获取设备列表
// yunos.tvpubadmin.device.yks.bots
//
// 获取设备列表
func Yunostvpubadmindeviceyksbots(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindeviceyksbotsAPIRequest, session string) (*tvupadmin.YunostvpubadmindeviceyksbotsAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindeviceyksbotsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
