package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindevicequery 获取设备列表
// yunos.tvpubadmin.device.query
//
// 获取设备列表
func Yunostvpubadmindevicequery(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindevicequeryAPIRequest, session string) (*tvupadmin.YunostvpubadmindevicequeryAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindevicequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
