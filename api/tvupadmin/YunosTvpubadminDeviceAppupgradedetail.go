package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindeviceappupgradedetail 获取应用升级详情
// yunos.tvpubadmin.device.appupgradedetail
//
// 获取应用升级详情
func Yunostvpubadmindeviceappupgradedetail(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindeviceappupgradedetailAPIRequest, session string) (*tvupadmin.YunostvpubadmindeviceappupgradedetailAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindeviceappupgradedetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
