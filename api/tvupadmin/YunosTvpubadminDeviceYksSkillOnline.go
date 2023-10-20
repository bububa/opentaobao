package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindeviceyksskillonline 迎客松技能上架接口
// yunos.tvpubadmin.device.yks.skill.online
//
// 迎客松技能上架接口
func Yunostvpubadmindeviceyksskillonline(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindeviceyksskillonlineAPIRequest, session string) (*tvupadmin.YunostvpubadmindeviceyksskillonlineAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindeviceyksskillonlineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
