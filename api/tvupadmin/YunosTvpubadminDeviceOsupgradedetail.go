package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceOsupgradedetail 获取系统升级详情
// yunos.tvpubadmin.device.osupgradedetail
//
// 获取系统升级详情
func YunosTvpubadminDeviceOsupgradedetail(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceOsupgradedetailAPIRequest, session string) (*tvupadmin.YunosTvpubadminDeviceOsupgradedetailAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminDeviceOsupgradedetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
