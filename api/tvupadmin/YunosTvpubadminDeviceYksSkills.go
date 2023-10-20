package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindeviceyksskills 根据设备id获取技能列表
// yunos.tvpubadmin.device.yks.skills
//
// 根据设备id获取技能列表
func Yunostvpubadmindeviceyksskills(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindeviceyksskillsAPIRequest, session string) (*tvupadmin.YunostvpubadmindeviceyksskillsAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindeviceyksskillsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
