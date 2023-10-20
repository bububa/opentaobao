package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindeviceosupgradedetail 获取系统升级详情
// yunos.tvpubadmin.device.osupgradedetail
//
// 获取系统升级详情
func Yunostvpubadmindeviceosupgradedetail(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindeviceosupgradedetailAPIRequest, session string) (*tvupadmin.YunostvpubadmindeviceosupgradedetailAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindeviceosupgradedetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
