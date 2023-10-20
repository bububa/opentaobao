package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindeviceyksskilloffline 技能下架
// yunos.tvpubadmin.device.yks.skill.offline
//
// 迎客松平台技能下架
func Yunostvpubadmindeviceyksskilloffline(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindeviceyksskillofflineAPIRequest, session string) (*tvupadmin.YunostvpubadmindeviceyksskillofflineAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindeviceyksskillofflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
