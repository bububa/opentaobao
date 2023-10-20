package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindeviceosupgradequery 系统升级查询
// yunos.tvpubadmin.device.osupgradequery
//
// 系统升级查询
func Yunostvpubadmindeviceosupgradequery(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindeviceosupgradequeryAPIRequest, session string) (*tvupadmin.YunostvpubadmindeviceosupgradequeryAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindeviceosupgradequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
