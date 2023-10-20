package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindeviceapkinfo 获取停开服apk信息
// yunos.tvpubadmin.device.apkinfo
//
// 获取停开服apk信息
func Yunostvpubadmindeviceapkinfo(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindeviceapkinfoAPIRequest, session string) (*tvupadmin.YunostvpubadmindeviceapkinfoAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindeviceapkinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
