package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindeviceapks 获取停开服apk列表
// yunos.tvpubadmin.device.apks
//
// 获取停开服apk列表
func Yunostvpubadmindeviceapks(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindeviceapksAPIRequest, session string) (*tvupadmin.YunostvpubadmindeviceapksAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindeviceapksAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
