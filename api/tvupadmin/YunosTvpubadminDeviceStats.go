package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindevicestats 获取设备统计数据
// yunos.tvpubadmin.device.stats
//
// 获取设备统计数据
func Yunostvpubadmindevicestats(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindevicestatsAPIRequest, session string) (*tvupadmin.YunostvpubadmindevicestatsAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindevicestatsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
