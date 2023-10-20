package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindeviceyksskilladd 添加技能
// yunos.tvpubadmin.device.yks.skill.add
//
// 添加技能
func Yunostvpubadmindeviceyksskilladd(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindeviceyksskilladdAPIRequest, session string) (*tvupadmin.YunostvpubadmindeviceyksskilladdAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindeviceyksskilladdAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
