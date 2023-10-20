package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindeviceappupgradequery 应用升级查询
// yunos.tvpubadmin.device.appupgradequery
//
// 应用升级查询
func Yunostvpubadmindeviceappupgradequery(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindeviceappupgradequeryAPIRequest, session string) (*tvupadmin.YunostvpubadmindeviceappupgradequeryAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindeviceappupgradequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
